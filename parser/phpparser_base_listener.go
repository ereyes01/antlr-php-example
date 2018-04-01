// Code generated from PhpParser.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // PhpParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePhpParserListener is a complete listener for a parse tree produced by PhpParser.
type BasePhpParserListener struct{}

var _ PhpParserListener = &BasePhpParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePhpParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePhpParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePhpParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePhpParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterHtmlDocument is called when production htmlDocument is entered.
func (s *BasePhpParserListener) EnterHtmlDocument(ctx *HtmlDocumentContext) {}

// ExitHtmlDocument is called when production htmlDocument is exited.
func (s *BasePhpParserListener) ExitHtmlDocument(ctx *HtmlDocumentContext) {}

// EnterHtmlElementOrPhpBlock is called when production htmlElementOrPhpBlock is entered.
func (s *BasePhpParserListener) EnterHtmlElementOrPhpBlock(ctx *HtmlElementOrPhpBlockContext) {}

// ExitHtmlElementOrPhpBlock is called when production htmlElementOrPhpBlock is exited.
func (s *BasePhpParserListener) ExitHtmlElementOrPhpBlock(ctx *HtmlElementOrPhpBlockContext) {}

// EnterHtmlElements is called when production htmlElements is entered.
func (s *BasePhpParserListener) EnterHtmlElements(ctx *HtmlElementsContext) {}

// ExitHtmlElements is called when production htmlElements is exited.
func (s *BasePhpParserListener) ExitHtmlElements(ctx *HtmlElementsContext) {}

// EnterHtmlElement is called when production htmlElement is entered.
func (s *BasePhpParserListener) EnterHtmlElement(ctx *HtmlElementContext) {}

// ExitHtmlElement is called when production htmlElement is exited.
func (s *BasePhpParserListener) ExitHtmlElement(ctx *HtmlElementContext) {}

// EnterScriptTextPart is called when production scriptTextPart is entered.
func (s *BasePhpParserListener) EnterScriptTextPart(ctx *ScriptTextPartContext) {}

// ExitScriptTextPart is called when production scriptTextPart is exited.
func (s *BasePhpParserListener) ExitScriptTextPart(ctx *ScriptTextPartContext) {}

// EnterPhpBlock is called when production phpBlock is entered.
func (s *BasePhpParserListener) EnterPhpBlock(ctx *PhpBlockContext) {}

// ExitPhpBlock is called when production phpBlock is exited.
func (s *BasePhpParserListener) ExitPhpBlock(ctx *PhpBlockContext) {}

// EnterImportStatement is called when production importStatement is entered.
func (s *BasePhpParserListener) EnterImportStatement(ctx *ImportStatementContext) {}

// ExitImportStatement is called when production importStatement is exited.
func (s *BasePhpParserListener) ExitImportStatement(ctx *ImportStatementContext) {}

// EnterTopStatement is called when production topStatement is entered.
func (s *BasePhpParserListener) EnterTopStatement(ctx *TopStatementContext) {}

// ExitTopStatement is called when production topStatement is exited.
func (s *BasePhpParserListener) ExitTopStatement(ctx *TopStatementContext) {}

// EnterUseDeclaration is called when production useDeclaration is entered.
func (s *BasePhpParserListener) EnterUseDeclaration(ctx *UseDeclarationContext) {}

// ExitUseDeclaration is called when production useDeclaration is exited.
func (s *BasePhpParserListener) ExitUseDeclaration(ctx *UseDeclarationContext) {}

// EnterUseDeclarationContentList is called when production useDeclarationContentList is entered.
func (s *BasePhpParserListener) EnterUseDeclarationContentList(ctx *UseDeclarationContentListContext) {
}

// ExitUseDeclarationContentList is called when production useDeclarationContentList is exited.
func (s *BasePhpParserListener) ExitUseDeclarationContentList(ctx *UseDeclarationContentListContext) {}

// EnterUseDeclarationContent is called when production useDeclarationContent is entered.
func (s *BasePhpParserListener) EnterUseDeclarationContent(ctx *UseDeclarationContentContext) {}

// ExitUseDeclarationContent is called when production useDeclarationContent is exited.
func (s *BasePhpParserListener) ExitUseDeclarationContent(ctx *UseDeclarationContentContext) {}

// EnterNamespaceDeclaration is called when production namespaceDeclaration is entered.
func (s *BasePhpParserListener) EnterNamespaceDeclaration(ctx *NamespaceDeclarationContext) {}

// ExitNamespaceDeclaration is called when production namespaceDeclaration is exited.
func (s *BasePhpParserListener) ExitNamespaceDeclaration(ctx *NamespaceDeclarationContext) {}

// EnterNamespaceStatement is called when production namespaceStatement is entered.
func (s *BasePhpParserListener) EnterNamespaceStatement(ctx *NamespaceStatementContext) {}

// ExitNamespaceStatement is called when production namespaceStatement is exited.
func (s *BasePhpParserListener) ExitNamespaceStatement(ctx *NamespaceStatementContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BasePhpParserListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BasePhpParserListener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterClassDeclaration is called when production classDeclaration is entered.
func (s *BasePhpParserListener) EnterClassDeclaration(ctx *ClassDeclarationContext) {}

// ExitClassDeclaration is called when production classDeclaration is exited.
func (s *BasePhpParserListener) ExitClassDeclaration(ctx *ClassDeclarationContext) {}

// EnterClassEntryType is called when production classEntryType is entered.
func (s *BasePhpParserListener) EnterClassEntryType(ctx *ClassEntryTypeContext) {}

// ExitClassEntryType is called when production classEntryType is exited.
func (s *BasePhpParserListener) ExitClassEntryType(ctx *ClassEntryTypeContext) {}

// EnterInterfaceList is called when production interfaceList is entered.
func (s *BasePhpParserListener) EnterInterfaceList(ctx *InterfaceListContext) {}

// ExitInterfaceList is called when production interfaceList is exited.
func (s *BasePhpParserListener) ExitInterfaceList(ctx *InterfaceListContext) {}

// EnterTypeParameterListInBrackets is called when production typeParameterListInBrackets is entered.
func (s *BasePhpParserListener) EnterTypeParameterListInBrackets(ctx *TypeParameterListInBracketsContext) {
}

// ExitTypeParameterListInBrackets is called when production typeParameterListInBrackets is exited.
func (s *BasePhpParserListener) ExitTypeParameterListInBrackets(ctx *TypeParameterListInBracketsContext) {
}

// EnterTypeParameterList is called when production typeParameterList is entered.
func (s *BasePhpParserListener) EnterTypeParameterList(ctx *TypeParameterListContext) {}

// ExitTypeParameterList is called when production typeParameterList is exited.
func (s *BasePhpParserListener) ExitTypeParameterList(ctx *TypeParameterListContext) {}

// EnterTypeParameterWithDefaultsList is called when production typeParameterWithDefaultsList is entered.
func (s *BasePhpParserListener) EnterTypeParameterWithDefaultsList(ctx *TypeParameterWithDefaultsListContext) {
}

// ExitTypeParameterWithDefaultsList is called when production typeParameterWithDefaultsList is exited.
func (s *BasePhpParserListener) ExitTypeParameterWithDefaultsList(ctx *TypeParameterWithDefaultsListContext) {
}

// EnterTypeParameterDecl is called when production typeParameterDecl is entered.
func (s *BasePhpParserListener) EnterTypeParameterDecl(ctx *TypeParameterDeclContext) {}

// ExitTypeParameterDecl is called when production typeParameterDecl is exited.
func (s *BasePhpParserListener) ExitTypeParameterDecl(ctx *TypeParameterDeclContext) {}

// EnterTypeParameterWithDefaultDecl is called when production typeParameterWithDefaultDecl is entered.
func (s *BasePhpParserListener) EnterTypeParameterWithDefaultDecl(ctx *TypeParameterWithDefaultDeclContext) {
}

// ExitTypeParameterWithDefaultDecl is called when production typeParameterWithDefaultDecl is exited.
func (s *BasePhpParserListener) ExitTypeParameterWithDefaultDecl(ctx *TypeParameterWithDefaultDeclContext) {
}

// EnterGenericDynamicArgs is called when production genericDynamicArgs is entered.
func (s *BasePhpParserListener) EnterGenericDynamicArgs(ctx *GenericDynamicArgsContext) {}

// ExitGenericDynamicArgs is called when production genericDynamicArgs is exited.
func (s *BasePhpParserListener) ExitGenericDynamicArgs(ctx *GenericDynamicArgsContext) {}

// EnterAttributes is called when production attributes is entered.
func (s *BasePhpParserListener) EnterAttributes(ctx *AttributesContext) {}

// ExitAttributes is called when production attributes is exited.
func (s *BasePhpParserListener) ExitAttributes(ctx *AttributesContext) {}

// EnterAttributesGroup is called when production attributesGroup is entered.
func (s *BasePhpParserListener) EnterAttributesGroup(ctx *AttributesGroupContext) {}

// ExitAttributesGroup is called when production attributesGroup is exited.
func (s *BasePhpParserListener) ExitAttributesGroup(ctx *AttributesGroupContext) {}

// EnterAttribute is called when production attribute is entered.
func (s *BasePhpParserListener) EnterAttribute(ctx *AttributeContext) {}

// ExitAttribute is called when production attribute is exited.
func (s *BasePhpParserListener) ExitAttribute(ctx *AttributeContext) {}

// EnterAttributeArgList is called when production attributeArgList is entered.
func (s *BasePhpParserListener) EnterAttributeArgList(ctx *AttributeArgListContext) {}

// ExitAttributeArgList is called when production attributeArgList is exited.
func (s *BasePhpParserListener) ExitAttributeArgList(ctx *AttributeArgListContext) {}

// EnterAttributeNamedArgList is called when production attributeNamedArgList is entered.
func (s *BasePhpParserListener) EnterAttributeNamedArgList(ctx *AttributeNamedArgListContext) {}

// ExitAttributeNamedArgList is called when production attributeNamedArgList is exited.
func (s *BasePhpParserListener) ExitAttributeNamedArgList(ctx *AttributeNamedArgListContext) {}

// EnterAttributeNamedArg is called when production attributeNamedArg is entered.
func (s *BasePhpParserListener) EnterAttributeNamedArg(ctx *AttributeNamedArgContext) {}

// ExitAttributeNamedArg is called when production attributeNamedArg is exited.
func (s *BasePhpParserListener) ExitAttributeNamedArg(ctx *AttributeNamedArgContext) {}

// EnterInnerStatementList is called when production innerStatementList is entered.
func (s *BasePhpParserListener) EnterInnerStatementList(ctx *InnerStatementListContext) {}

// ExitInnerStatementList is called when production innerStatementList is exited.
func (s *BasePhpParserListener) ExitInnerStatementList(ctx *InnerStatementListContext) {}

// EnterInnerStatement is called when production innerStatement is entered.
func (s *BasePhpParserListener) EnterInnerStatement(ctx *InnerStatementContext) {}

// ExitInnerStatement is called when production innerStatement is exited.
func (s *BasePhpParserListener) ExitInnerStatement(ctx *InnerStatementContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasePhpParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasePhpParserListener) ExitStatement(ctx *StatementContext) {}

// EnterVoidStatement is called when production voidStatement is entered.
func (s *BasePhpParserListener) EnterVoidStatement(ctx *VoidStatementContext) {}

// ExitVoidStatement is called when production voidStatement is exited.
func (s *BasePhpParserListener) ExitVoidStatement(ctx *VoidStatementContext) {}

// EnterBlockStatement is called when production blockStatement is entered.
func (s *BasePhpParserListener) EnterBlockStatement(ctx *BlockStatementContext) {}

// ExitBlockStatement is called when production blockStatement is exited.
func (s *BasePhpParserListener) ExitBlockStatement(ctx *BlockStatementContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BasePhpParserListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BasePhpParserListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterElseIfStatement is called when production elseIfStatement is entered.
func (s *BasePhpParserListener) EnterElseIfStatement(ctx *ElseIfStatementContext) {}

// ExitElseIfStatement is called when production elseIfStatement is exited.
func (s *BasePhpParserListener) ExitElseIfStatement(ctx *ElseIfStatementContext) {}

// EnterElseIfColonStatement is called when production elseIfColonStatement is entered.
func (s *BasePhpParserListener) EnterElseIfColonStatement(ctx *ElseIfColonStatementContext) {}

// ExitElseIfColonStatement is called when production elseIfColonStatement is exited.
func (s *BasePhpParserListener) ExitElseIfColonStatement(ctx *ElseIfColonStatementContext) {}

// EnterElseStatement is called when production elseStatement is entered.
func (s *BasePhpParserListener) EnterElseStatement(ctx *ElseStatementContext) {}

// ExitElseStatement is called when production elseStatement is exited.
func (s *BasePhpParserListener) ExitElseStatement(ctx *ElseStatementContext) {}

// EnterElseColonStatement is called when production elseColonStatement is entered.
func (s *BasePhpParserListener) EnterElseColonStatement(ctx *ElseColonStatementContext) {}

// ExitElseColonStatement is called when production elseColonStatement is exited.
func (s *BasePhpParserListener) ExitElseColonStatement(ctx *ElseColonStatementContext) {}

// EnterWhileStatement is called when production whileStatement is entered.
func (s *BasePhpParserListener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production whileStatement is exited.
func (s *BasePhpParserListener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterDoWhileStatement is called when production doWhileStatement is entered.
func (s *BasePhpParserListener) EnterDoWhileStatement(ctx *DoWhileStatementContext) {}

// ExitDoWhileStatement is called when production doWhileStatement is exited.
func (s *BasePhpParserListener) ExitDoWhileStatement(ctx *DoWhileStatementContext) {}

// EnterForStatement is called when production forStatement is entered.
func (s *BasePhpParserListener) EnterForStatement(ctx *ForStatementContext) {}

// ExitForStatement is called when production forStatement is exited.
func (s *BasePhpParserListener) ExitForStatement(ctx *ForStatementContext) {}

// EnterForInit is called when production forInit is entered.
func (s *BasePhpParserListener) EnterForInit(ctx *ForInitContext) {}

// ExitForInit is called when production forInit is exited.
func (s *BasePhpParserListener) ExitForInit(ctx *ForInitContext) {}

// EnterForUpdate is called when production forUpdate is entered.
func (s *BasePhpParserListener) EnterForUpdate(ctx *ForUpdateContext) {}

// ExitForUpdate is called when production forUpdate is exited.
func (s *BasePhpParserListener) ExitForUpdate(ctx *ForUpdateContext) {}

// EnterSwitchStatement is called when production switchStatement is entered.
func (s *BasePhpParserListener) EnterSwitchStatement(ctx *SwitchStatementContext) {}

// ExitSwitchStatement is called when production switchStatement is exited.
func (s *BasePhpParserListener) ExitSwitchStatement(ctx *SwitchStatementContext) {}

// EnterSwitchBlock is called when production switchBlock is entered.
func (s *BasePhpParserListener) EnterSwitchBlock(ctx *SwitchBlockContext) {}

// ExitSwitchBlock is called when production switchBlock is exited.
func (s *BasePhpParserListener) ExitSwitchBlock(ctx *SwitchBlockContext) {}

// EnterBreakStatement is called when production breakStatement is entered.
func (s *BasePhpParserListener) EnterBreakStatement(ctx *BreakStatementContext) {}

// ExitBreakStatement is called when production breakStatement is exited.
func (s *BasePhpParserListener) ExitBreakStatement(ctx *BreakStatementContext) {}

// EnterContinueStatement is called when production continueStatement is entered.
func (s *BasePhpParserListener) EnterContinueStatement(ctx *ContinueStatementContext) {}

// ExitContinueStatement is called when production continueStatement is exited.
func (s *BasePhpParserListener) ExitContinueStatement(ctx *ContinueStatementContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BasePhpParserListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BasePhpParserListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterExpressionStatement is called when production expressionStatement is entered.
func (s *BasePhpParserListener) EnterExpressionStatement(ctx *ExpressionStatementContext) {}

// ExitExpressionStatement is called when production expressionStatement is exited.
func (s *BasePhpParserListener) ExitExpressionStatement(ctx *ExpressionStatementContext) {}

// EnterUnsetStatement is called when production unsetStatement is entered.
func (s *BasePhpParserListener) EnterUnsetStatement(ctx *UnsetStatementContext) {}

// ExitUnsetStatement is called when production unsetStatement is exited.
func (s *BasePhpParserListener) ExitUnsetStatement(ctx *UnsetStatementContext) {}

// EnterForeachStatement is called when production foreachStatement is entered.
func (s *BasePhpParserListener) EnterForeachStatement(ctx *ForeachStatementContext) {}

// ExitForeachStatement is called when production foreachStatement is exited.
func (s *BasePhpParserListener) ExitForeachStatement(ctx *ForeachStatementContext) {}

// EnterTryCatchFinally is called when production tryCatchFinally is entered.
func (s *BasePhpParserListener) EnterTryCatchFinally(ctx *TryCatchFinallyContext) {}

// ExitTryCatchFinally is called when production tryCatchFinally is exited.
func (s *BasePhpParserListener) ExitTryCatchFinally(ctx *TryCatchFinallyContext) {}

// EnterCatchClause is called when production catchClause is entered.
func (s *BasePhpParserListener) EnterCatchClause(ctx *CatchClauseContext) {}

// ExitCatchClause is called when production catchClause is exited.
func (s *BasePhpParserListener) ExitCatchClause(ctx *CatchClauseContext) {}

// EnterFinallyStatement is called when production finallyStatement is entered.
func (s *BasePhpParserListener) EnterFinallyStatement(ctx *FinallyStatementContext) {}

// ExitFinallyStatement is called when production finallyStatement is exited.
func (s *BasePhpParserListener) ExitFinallyStatement(ctx *FinallyStatementContext) {}

// EnterThrowStatement is called when production throwStatement is entered.
func (s *BasePhpParserListener) EnterThrowStatement(ctx *ThrowStatementContext) {}

// ExitThrowStatement is called when production throwStatement is exited.
func (s *BasePhpParserListener) ExitThrowStatement(ctx *ThrowStatementContext) {}

// EnterGotoStatement is called when production gotoStatement is entered.
func (s *BasePhpParserListener) EnterGotoStatement(ctx *GotoStatementContext) {}

// ExitGotoStatement is called when production gotoStatement is exited.
func (s *BasePhpParserListener) ExitGotoStatement(ctx *GotoStatementContext) {}

// EnterDeclareStatement is called when production declareStatement is entered.
func (s *BasePhpParserListener) EnterDeclareStatement(ctx *DeclareStatementContext) {}

// ExitDeclareStatement is called when production declareStatement is exited.
func (s *BasePhpParserListener) ExitDeclareStatement(ctx *DeclareStatementContext) {}

// EnterInlineHtmlStatement is called when production inlineHtmlStatement is entered.
func (s *BasePhpParserListener) EnterInlineHtmlStatement(ctx *InlineHtmlStatementContext) {}

// ExitInlineHtmlStatement is called when production inlineHtmlStatement is exited.
func (s *BasePhpParserListener) ExitInlineHtmlStatement(ctx *InlineHtmlStatementContext) {}

// EnterInlineHtml is called when production inlineHtml is entered.
func (s *BasePhpParserListener) EnterInlineHtml(ctx *InlineHtmlContext) {}

// ExitInlineHtml is called when production inlineHtml is exited.
func (s *BasePhpParserListener) ExitInlineHtml(ctx *InlineHtmlContext) {}

// EnterDeclareList is called when production declareList is entered.
func (s *BasePhpParserListener) EnterDeclareList(ctx *DeclareListContext) {}

// ExitDeclareList is called when production declareList is exited.
func (s *BasePhpParserListener) ExitDeclareList(ctx *DeclareListContext) {}

// EnterFormalParameterList is called when production formalParameterList is entered.
func (s *BasePhpParserListener) EnterFormalParameterList(ctx *FormalParameterListContext) {}

// ExitFormalParameterList is called when production formalParameterList is exited.
func (s *BasePhpParserListener) ExitFormalParameterList(ctx *FormalParameterListContext) {}

// EnterFormalParameter is called when production formalParameter is entered.
func (s *BasePhpParserListener) EnterFormalParameter(ctx *FormalParameterContext) {}

// ExitFormalParameter is called when production formalParameter is exited.
func (s *BasePhpParserListener) ExitFormalParameter(ctx *FormalParameterContext) {}

// EnterTypeHint is called when production typeHint is entered.
func (s *BasePhpParserListener) EnterTypeHint(ctx *TypeHintContext) {}

// ExitTypeHint is called when production typeHint is exited.
func (s *BasePhpParserListener) ExitTypeHint(ctx *TypeHintContext) {}

// EnterGlobalStatement is called when production globalStatement is entered.
func (s *BasePhpParserListener) EnterGlobalStatement(ctx *GlobalStatementContext) {}

// ExitGlobalStatement is called when production globalStatement is exited.
func (s *BasePhpParserListener) ExitGlobalStatement(ctx *GlobalStatementContext) {}

// EnterGlobalVar is called when production globalVar is entered.
func (s *BasePhpParserListener) EnterGlobalVar(ctx *GlobalVarContext) {}

// ExitGlobalVar is called when production globalVar is exited.
func (s *BasePhpParserListener) ExitGlobalVar(ctx *GlobalVarContext) {}

// EnterEchoStatement is called when production echoStatement is entered.
func (s *BasePhpParserListener) EnterEchoStatement(ctx *EchoStatementContext) {}

// ExitEchoStatement is called when production echoStatement is exited.
func (s *BasePhpParserListener) ExitEchoStatement(ctx *EchoStatementContext) {}

// EnterStaticVariableStatement is called when production staticVariableStatement is entered.
func (s *BasePhpParserListener) EnterStaticVariableStatement(ctx *StaticVariableStatementContext) {}

// ExitStaticVariableStatement is called when production staticVariableStatement is exited.
func (s *BasePhpParserListener) ExitStaticVariableStatement(ctx *StaticVariableStatementContext) {}

// EnterClassStatement is called when production classStatement is entered.
func (s *BasePhpParserListener) EnterClassStatement(ctx *ClassStatementContext) {}

// ExitClassStatement is called when production classStatement is exited.
func (s *BasePhpParserListener) ExitClassStatement(ctx *ClassStatementContext) {}

// EnterTraitAdaptations is called when production traitAdaptations is entered.
func (s *BasePhpParserListener) EnterTraitAdaptations(ctx *TraitAdaptationsContext) {}

// ExitTraitAdaptations is called when production traitAdaptations is exited.
func (s *BasePhpParserListener) ExitTraitAdaptations(ctx *TraitAdaptationsContext) {}

// EnterTraitAdaptationStatement is called when production traitAdaptationStatement is entered.
func (s *BasePhpParserListener) EnterTraitAdaptationStatement(ctx *TraitAdaptationStatementContext) {}

// ExitTraitAdaptationStatement is called when production traitAdaptationStatement is exited.
func (s *BasePhpParserListener) ExitTraitAdaptationStatement(ctx *TraitAdaptationStatementContext) {}

// EnterTraitPrecedence is called when production traitPrecedence is entered.
func (s *BasePhpParserListener) EnterTraitPrecedence(ctx *TraitPrecedenceContext) {}

// ExitTraitPrecedence is called when production traitPrecedence is exited.
func (s *BasePhpParserListener) ExitTraitPrecedence(ctx *TraitPrecedenceContext) {}

// EnterTraitAlias is called when production traitAlias is entered.
func (s *BasePhpParserListener) EnterTraitAlias(ctx *TraitAliasContext) {}

// ExitTraitAlias is called when production traitAlias is exited.
func (s *BasePhpParserListener) ExitTraitAlias(ctx *TraitAliasContext) {}

// EnterTraitMethodReference is called when production traitMethodReference is entered.
func (s *BasePhpParserListener) EnterTraitMethodReference(ctx *TraitMethodReferenceContext) {}

// ExitTraitMethodReference is called when production traitMethodReference is exited.
func (s *BasePhpParserListener) ExitTraitMethodReference(ctx *TraitMethodReferenceContext) {}

// EnterBaseCtorCall is called when production baseCtorCall is entered.
func (s *BasePhpParserListener) EnterBaseCtorCall(ctx *BaseCtorCallContext) {}

// ExitBaseCtorCall is called when production baseCtorCall is exited.
func (s *BasePhpParserListener) ExitBaseCtorCall(ctx *BaseCtorCallContext) {}

// EnterMethodBody is called when production methodBody is entered.
func (s *BasePhpParserListener) EnterMethodBody(ctx *MethodBodyContext) {}

// ExitMethodBody is called when production methodBody is exited.
func (s *BasePhpParserListener) ExitMethodBody(ctx *MethodBodyContext) {}

// EnterPropertyModifiers is called when production propertyModifiers is entered.
func (s *BasePhpParserListener) EnterPropertyModifiers(ctx *PropertyModifiersContext) {}

// ExitPropertyModifiers is called when production propertyModifiers is exited.
func (s *BasePhpParserListener) ExitPropertyModifiers(ctx *PropertyModifiersContext) {}

// EnterMemberModifiers is called when production memberModifiers is entered.
func (s *BasePhpParserListener) EnterMemberModifiers(ctx *MemberModifiersContext) {}

// ExitMemberModifiers is called when production memberModifiers is exited.
func (s *BasePhpParserListener) ExitMemberModifiers(ctx *MemberModifiersContext) {}

// EnterVariableInitializer is called when production variableInitializer is entered.
func (s *BasePhpParserListener) EnterVariableInitializer(ctx *VariableInitializerContext) {}

// ExitVariableInitializer is called when production variableInitializer is exited.
func (s *BasePhpParserListener) ExitVariableInitializer(ctx *VariableInitializerContext) {}

// EnterIdentifierInititalizer is called when production identifierInititalizer is entered.
func (s *BasePhpParserListener) EnterIdentifierInititalizer(ctx *IdentifierInititalizerContext) {}

// ExitIdentifierInititalizer is called when production identifierInititalizer is exited.
func (s *BasePhpParserListener) ExitIdentifierInititalizer(ctx *IdentifierInititalizerContext) {}

// EnterGlobalConstantDeclaration is called when production globalConstantDeclaration is entered.
func (s *BasePhpParserListener) EnterGlobalConstantDeclaration(ctx *GlobalConstantDeclarationContext) {
}

// ExitGlobalConstantDeclaration is called when production globalConstantDeclaration is exited.
func (s *BasePhpParserListener) ExitGlobalConstantDeclaration(ctx *GlobalConstantDeclarationContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BasePhpParserListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BasePhpParserListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterParenthesis is called when production parenthesis is entered.
func (s *BasePhpParserListener) EnterParenthesis(ctx *ParenthesisContext) {}

// ExitParenthesis is called when production parenthesis is exited.
func (s *BasePhpParserListener) ExitParenthesis(ctx *ParenthesisContext) {}

// EnterChainExpression is called when production ChainExpression is entered.
func (s *BasePhpParserListener) EnterChainExpression(ctx *ChainExpressionContext) {}

// ExitChainExpression is called when production ChainExpression is exited.
func (s *BasePhpParserListener) ExitChainExpression(ctx *ChainExpressionContext) {}

// EnterUnaryOperatorExpression is called when production UnaryOperatorExpression is entered.
func (s *BasePhpParserListener) EnterUnaryOperatorExpression(ctx *UnaryOperatorExpressionContext) {}

// ExitUnaryOperatorExpression is called when production UnaryOperatorExpression is exited.
func (s *BasePhpParserListener) ExitUnaryOperatorExpression(ctx *UnaryOperatorExpressionContext) {}

// EnterSpecialWordExpression is called when production SpecialWordExpression is entered.
func (s *BasePhpParserListener) EnterSpecialWordExpression(ctx *SpecialWordExpressionContext) {}

// ExitSpecialWordExpression is called when production SpecialWordExpression is exited.
func (s *BasePhpParserListener) ExitSpecialWordExpression(ctx *SpecialWordExpressionContext) {}

// EnterArrayCreationExpression is called when production ArrayCreationExpression is entered.
func (s *BasePhpParserListener) EnterArrayCreationExpression(ctx *ArrayCreationExpressionContext) {}

// ExitArrayCreationExpression is called when production ArrayCreationExpression is exited.
func (s *BasePhpParserListener) ExitArrayCreationExpression(ctx *ArrayCreationExpressionContext) {}

// EnterParenthesisExpression is called when production ParenthesisExpression is entered.
func (s *BasePhpParserListener) EnterParenthesisExpression(ctx *ParenthesisExpressionContext) {}

// ExitParenthesisExpression is called when production ParenthesisExpression is exited.
func (s *BasePhpParserListener) ExitParenthesisExpression(ctx *ParenthesisExpressionContext) {}

// EnterBackQuoteStringExpression is called when production BackQuoteStringExpression is entered.
func (s *BasePhpParserListener) EnterBackQuoteStringExpression(ctx *BackQuoteStringExpressionContext) {
}

// ExitBackQuoteStringExpression is called when production BackQuoteStringExpression is exited.
func (s *BasePhpParserListener) ExitBackQuoteStringExpression(ctx *BackQuoteStringExpressionContext) {}

// EnterConditionalExpression is called when production ConditionalExpression is entered.
func (s *BasePhpParserListener) EnterConditionalExpression(ctx *ConditionalExpressionContext) {}

// ExitConditionalExpression is called when production ConditionalExpression is exited.
func (s *BasePhpParserListener) ExitConditionalExpression(ctx *ConditionalExpressionContext) {}

// EnterNewExprClause is called when production NewExprClause is entered.
func (s *BasePhpParserListener) EnterNewExprClause(ctx *NewExprClauseContext) {}

// ExitNewExprClause is called when production NewExprClause is exited.
func (s *BasePhpParserListener) ExitNewExprClause(ctx *NewExprClauseContext) {}

// EnterArithmeticExpression is called when production ArithmeticExpression is entered.
func (s *BasePhpParserListener) EnterArithmeticExpression(ctx *ArithmeticExpressionContext) {}

// ExitArithmeticExpression is called when production ArithmeticExpression is exited.
func (s *BasePhpParserListener) ExitArithmeticExpression(ctx *ArithmeticExpressionContext) {}

// EnterIndexerExpression is called when production IndexerExpression is entered.
func (s *BasePhpParserListener) EnterIndexerExpression(ctx *IndexerExpressionContext) {}

// ExitIndexerExpression is called when production IndexerExpression is exited.
func (s *BasePhpParserListener) ExitIndexerExpression(ctx *IndexerExpressionContext) {}

// EnterScalarExpression is called when production ScalarExpression is entered.
func (s *BasePhpParserListener) EnterScalarExpression(ctx *ScalarExpressionContext) {}

// ExitScalarExpression is called when production ScalarExpression is exited.
func (s *BasePhpParserListener) ExitScalarExpression(ctx *ScalarExpressionContext) {}

// EnterPrefixIncDecExpression is called when production PrefixIncDecExpression is entered.
func (s *BasePhpParserListener) EnterPrefixIncDecExpression(ctx *PrefixIncDecExpressionContext) {}

// ExitPrefixIncDecExpression is called when production PrefixIncDecExpression is exited.
func (s *BasePhpParserListener) ExitPrefixIncDecExpression(ctx *PrefixIncDecExpressionContext) {}

// EnterComparisonExpression is called when production ComparisonExpression is entered.
func (s *BasePhpParserListener) EnterComparisonExpression(ctx *ComparisonExpressionContext) {}

// ExitComparisonExpression is called when production ComparisonExpression is exited.
func (s *BasePhpParserListener) ExitComparisonExpression(ctx *ComparisonExpressionContext) {}

// EnterLogicalExpression is called when production LogicalExpression is entered.
func (s *BasePhpParserListener) EnterLogicalExpression(ctx *LogicalExpressionContext) {}

// ExitLogicalExpression is called when production LogicalExpression is exited.
func (s *BasePhpParserListener) ExitLogicalExpression(ctx *LogicalExpressionContext) {}

// EnterPrintExpression is called when production PrintExpression is entered.
func (s *BasePhpParserListener) EnterPrintExpression(ctx *PrintExpressionContext) {}

// ExitPrintExpression is called when production PrintExpression is exited.
func (s *BasePhpParserListener) ExitPrintExpression(ctx *PrintExpressionContext) {}

// EnterAssignmentExpression is called when production AssignmentExpression is entered.
func (s *BasePhpParserListener) EnterAssignmentExpression(ctx *AssignmentExpressionContext) {}

// ExitAssignmentExpression is called when production AssignmentExpression is exited.
func (s *BasePhpParserListener) ExitAssignmentExpression(ctx *AssignmentExpressionContext) {}

// EnterPostfixIncDecExpression is called when production PostfixIncDecExpression is entered.
func (s *BasePhpParserListener) EnterPostfixIncDecExpression(ctx *PostfixIncDecExpressionContext) {}

// ExitPostfixIncDecExpression is called when production PostfixIncDecExpression is exited.
func (s *BasePhpParserListener) ExitPostfixIncDecExpression(ctx *PostfixIncDecExpressionContext) {}

// EnterCastExpression is called when production CastExpression is entered.
func (s *BasePhpParserListener) EnterCastExpression(ctx *CastExpressionContext) {}

// ExitCastExpression is called when production CastExpression is exited.
func (s *BasePhpParserListener) ExitCastExpression(ctx *CastExpressionContext) {}

// EnterInstanceOfExpression is called when production InstanceOfExpression is entered.
func (s *BasePhpParserListener) EnterInstanceOfExpression(ctx *InstanceOfExpressionContext) {}

// ExitInstanceOfExpression is called when production InstanceOfExpression is exited.
func (s *BasePhpParserListener) ExitInstanceOfExpression(ctx *InstanceOfExpressionContext) {}

// EnterLambdaFunctionExpression is called when production LambdaFunctionExpression is entered.
func (s *BasePhpParserListener) EnterLambdaFunctionExpression(ctx *LambdaFunctionExpressionContext) {}

// ExitLambdaFunctionExpression is called when production LambdaFunctionExpression is exited.
func (s *BasePhpParserListener) ExitLambdaFunctionExpression(ctx *LambdaFunctionExpressionContext) {}

// EnterBitwiseExpression is called when production BitwiseExpression is entered.
func (s *BasePhpParserListener) EnterBitwiseExpression(ctx *BitwiseExpressionContext) {}

// ExitBitwiseExpression is called when production BitwiseExpression is exited.
func (s *BasePhpParserListener) ExitBitwiseExpression(ctx *BitwiseExpressionContext) {}

// EnterCloneExpression is called when production CloneExpression is entered.
func (s *BasePhpParserListener) EnterCloneExpression(ctx *CloneExpressionContext) {}

// ExitCloneExpression is called when production CloneExpression is exited.
func (s *BasePhpParserListener) ExitCloneExpression(ctx *CloneExpressionContext) {}

// EnterNewExpr is called when production newExpr is entered.
func (s *BasePhpParserListener) EnterNewExpr(ctx *NewExprContext) {}

// ExitNewExpr is called when production newExpr is exited.
func (s *BasePhpParserListener) ExitNewExpr(ctx *NewExprContext) {}

// EnterAssignmentOperator is called when production assignmentOperator is entered.
func (s *BasePhpParserListener) EnterAssignmentOperator(ctx *AssignmentOperatorContext) {}

// ExitAssignmentOperator is called when production assignmentOperator is exited.
func (s *BasePhpParserListener) ExitAssignmentOperator(ctx *AssignmentOperatorContext) {}

// EnterYieldExpression is called when production yieldExpression is entered.
func (s *BasePhpParserListener) EnterYieldExpression(ctx *YieldExpressionContext) {}

// ExitYieldExpression is called when production yieldExpression is exited.
func (s *BasePhpParserListener) ExitYieldExpression(ctx *YieldExpressionContext) {}

// EnterArrayItemList is called when production arrayItemList is entered.
func (s *BasePhpParserListener) EnterArrayItemList(ctx *ArrayItemListContext) {}

// ExitArrayItemList is called when production arrayItemList is exited.
func (s *BasePhpParserListener) ExitArrayItemList(ctx *ArrayItemListContext) {}

// EnterArrayItem is called when production arrayItem is entered.
func (s *BasePhpParserListener) EnterArrayItem(ctx *ArrayItemContext) {}

// ExitArrayItem is called when production arrayItem is exited.
func (s *BasePhpParserListener) ExitArrayItem(ctx *ArrayItemContext) {}

// EnterLambdaFunctionUseVars is called when production lambdaFunctionUseVars is entered.
func (s *BasePhpParserListener) EnterLambdaFunctionUseVars(ctx *LambdaFunctionUseVarsContext) {}

// ExitLambdaFunctionUseVars is called when production lambdaFunctionUseVars is exited.
func (s *BasePhpParserListener) ExitLambdaFunctionUseVars(ctx *LambdaFunctionUseVarsContext) {}

// EnterLambdaFunctionUseVar is called when production lambdaFunctionUseVar is entered.
func (s *BasePhpParserListener) EnterLambdaFunctionUseVar(ctx *LambdaFunctionUseVarContext) {}

// ExitLambdaFunctionUseVar is called when production lambdaFunctionUseVar is exited.
func (s *BasePhpParserListener) ExitLambdaFunctionUseVar(ctx *LambdaFunctionUseVarContext) {}

// EnterQualifiedStaticTypeRef is called when production qualifiedStaticTypeRef is entered.
func (s *BasePhpParserListener) EnterQualifiedStaticTypeRef(ctx *QualifiedStaticTypeRefContext) {}

// ExitQualifiedStaticTypeRef is called when production qualifiedStaticTypeRef is exited.
func (s *BasePhpParserListener) ExitQualifiedStaticTypeRef(ctx *QualifiedStaticTypeRefContext) {}

// EnterTypeRef is called when production typeRef is entered.
func (s *BasePhpParserListener) EnterTypeRef(ctx *TypeRefContext) {}

// ExitTypeRef is called when production typeRef is exited.
func (s *BasePhpParserListener) ExitTypeRef(ctx *TypeRefContext) {}

// EnterIndirectTypeRef is called when production indirectTypeRef is entered.
func (s *BasePhpParserListener) EnterIndirectTypeRef(ctx *IndirectTypeRefContext) {}

// ExitIndirectTypeRef is called when production indirectTypeRef is exited.
func (s *BasePhpParserListener) ExitIndirectTypeRef(ctx *IndirectTypeRefContext) {}

// EnterQualifiedNamespaceName is called when production qualifiedNamespaceName is entered.
func (s *BasePhpParserListener) EnterQualifiedNamespaceName(ctx *QualifiedNamespaceNameContext) {}

// ExitQualifiedNamespaceName is called when production qualifiedNamespaceName is exited.
func (s *BasePhpParserListener) ExitQualifiedNamespaceName(ctx *QualifiedNamespaceNameContext) {}

// EnterNamespaceNameList is called when production namespaceNameList is entered.
func (s *BasePhpParserListener) EnterNamespaceNameList(ctx *NamespaceNameListContext) {}

// ExitNamespaceNameList is called when production namespaceNameList is exited.
func (s *BasePhpParserListener) ExitNamespaceNameList(ctx *NamespaceNameListContext) {}

// EnterQualifiedNamespaceNameList is called when production qualifiedNamespaceNameList is entered.
func (s *BasePhpParserListener) EnterQualifiedNamespaceNameList(ctx *QualifiedNamespaceNameListContext) {
}

// ExitQualifiedNamespaceNameList is called when production qualifiedNamespaceNameList is exited.
func (s *BasePhpParserListener) ExitQualifiedNamespaceNameList(ctx *QualifiedNamespaceNameListContext) {
}

// EnterArguments is called when production arguments is entered.
func (s *BasePhpParserListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BasePhpParserListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterActualArgument is called when production actualArgument is entered.
func (s *BasePhpParserListener) EnterActualArgument(ctx *ActualArgumentContext) {}

// ExitActualArgument is called when production actualArgument is exited.
func (s *BasePhpParserListener) ExitActualArgument(ctx *ActualArgumentContext) {}

// EnterConstantInititalizer is called when production constantInititalizer is entered.
func (s *BasePhpParserListener) EnterConstantInititalizer(ctx *ConstantInititalizerContext) {}

// ExitConstantInititalizer is called when production constantInititalizer is exited.
func (s *BasePhpParserListener) ExitConstantInititalizer(ctx *ConstantInititalizerContext) {}

// EnterConstantArrayItemList is called when production constantArrayItemList is entered.
func (s *BasePhpParserListener) EnterConstantArrayItemList(ctx *ConstantArrayItemListContext) {}

// ExitConstantArrayItemList is called when production constantArrayItemList is exited.
func (s *BasePhpParserListener) ExitConstantArrayItemList(ctx *ConstantArrayItemListContext) {}

// EnterConstantArrayItem is called when production constantArrayItem is entered.
func (s *BasePhpParserListener) EnterConstantArrayItem(ctx *ConstantArrayItemContext) {}

// ExitConstantArrayItem is called when production constantArrayItem is exited.
func (s *BasePhpParserListener) ExitConstantArrayItem(ctx *ConstantArrayItemContext) {}

// EnterConstant is called when production constant is entered.
func (s *BasePhpParserListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BasePhpParserListener) ExitConstant(ctx *ConstantContext) {}

// EnterLiteralConstant is called when production literalConstant is entered.
func (s *BasePhpParserListener) EnterLiteralConstant(ctx *LiteralConstantContext) {}

// ExitLiteralConstant is called when production literalConstant is exited.
func (s *BasePhpParserListener) ExitLiteralConstant(ctx *LiteralConstantContext) {}

// EnterNumericConstant is called when production numericConstant is entered.
func (s *BasePhpParserListener) EnterNumericConstant(ctx *NumericConstantContext) {}

// ExitNumericConstant is called when production numericConstant is exited.
func (s *BasePhpParserListener) ExitNumericConstant(ctx *NumericConstantContext) {}

// EnterClassConstant is called when production classConstant is entered.
func (s *BasePhpParserListener) EnterClassConstant(ctx *ClassConstantContext) {}

// ExitClassConstant is called when production classConstant is exited.
func (s *BasePhpParserListener) ExitClassConstant(ctx *ClassConstantContext) {}

// EnterStringConstant is called when production stringConstant is entered.
func (s *BasePhpParserListener) EnterStringConstant(ctx *StringConstantContext) {}

// ExitStringConstant is called when production stringConstant is exited.
func (s *BasePhpParserListener) ExitStringConstant(ctx *StringConstantContext) {}

// EnterStr is called when production str is entered.
func (s *BasePhpParserListener) EnterStr(ctx *StrContext) {}

// ExitStr is called when production str is exited.
func (s *BasePhpParserListener) ExitStr(ctx *StrContext) {}

// EnterInterpolatedStringPart is called when production interpolatedStringPart is entered.
func (s *BasePhpParserListener) EnterInterpolatedStringPart(ctx *InterpolatedStringPartContext) {}

// ExitInterpolatedStringPart is called when production interpolatedStringPart is exited.
func (s *BasePhpParserListener) ExitInterpolatedStringPart(ctx *InterpolatedStringPartContext) {}

// EnterChainList is called when production chainList is entered.
func (s *BasePhpParserListener) EnterChainList(ctx *ChainListContext) {}

// ExitChainList is called when production chainList is exited.
func (s *BasePhpParserListener) ExitChainList(ctx *ChainListContext) {}

// EnterChain is called when production chain is entered.
func (s *BasePhpParserListener) EnterChain(ctx *ChainContext) {}

// ExitChain is called when production chain is exited.
func (s *BasePhpParserListener) ExitChain(ctx *ChainContext) {}

// EnterMemberAccess is called when production memberAccess is entered.
func (s *BasePhpParserListener) EnterMemberAccess(ctx *MemberAccessContext) {}

// ExitMemberAccess is called when production memberAccess is exited.
func (s *BasePhpParserListener) ExitMemberAccess(ctx *MemberAccessContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BasePhpParserListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BasePhpParserListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterFunctionCallName is called when production functionCallName is entered.
func (s *BasePhpParserListener) EnterFunctionCallName(ctx *FunctionCallNameContext) {}

// ExitFunctionCallName is called when production functionCallName is exited.
func (s *BasePhpParserListener) ExitFunctionCallName(ctx *FunctionCallNameContext) {}

// EnterActualArguments is called when production actualArguments is entered.
func (s *BasePhpParserListener) EnterActualArguments(ctx *ActualArgumentsContext) {}

// ExitActualArguments is called when production actualArguments is exited.
func (s *BasePhpParserListener) ExitActualArguments(ctx *ActualArgumentsContext) {}

// EnterChainBase is called when production chainBase is entered.
func (s *BasePhpParserListener) EnterChainBase(ctx *ChainBaseContext) {}

// ExitChainBase is called when production chainBase is exited.
func (s *BasePhpParserListener) ExitChainBase(ctx *ChainBaseContext) {}

// EnterKeyedFieldName is called when production keyedFieldName is entered.
func (s *BasePhpParserListener) EnterKeyedFieldName(ctx *KeyedFieldNameContext) {}

// ExitKeyedFieldName is called when production keyedFieldName is exited.
func (s *BasePhpParserListener) ExitKeyedFieldName(ctx *KeyedFieldNameContext) {}

// EnterKeyedSimpleFieldName is called when production keyedSimpleFieldName is entered.
func (s *BasePhpParserListener) EnterKeyedSimpleFieldName(ctx *KeyedSimpleFieldNameContext) {}

// ExitKeyedSimpleFieldName is called when production keyedSimpleFieldName is exited.
func (s *BasePhpParserListener) ExitKeyedSimpleFieldName(ctx *KeyedSimpleFieldNameContext) {}

// EnterKeyedVariable is called when production keyedVariable is entered.
func (s *BasePhpParserListener) EnterKeyedVariable(ctx *KeyedVariableContext) {}

// ExitKeyedVariable is called when production keyedVariable is exited.
func (s *BasePhpParserListener) ExitKeyedVariable(ctx *KeyedVariableContext) {}

// EnterSquareCurlyExpression is called when production squareCurlyExpression is entered.
func (s *BasePhpParserListener) EnterSquareCurlyExpression(ctx *SquareCurlyExpressionContext) {}

// ExitSquareCurlyExpression is called when production squareCurlyExpression is exited.
func (s *BasePhpParserListener) ExitSquareCurlyExpression(ctx *SquareCurlyExpressionContext) {}

// EnterAssignmentList is called when production assignmentList is entered.
func (s *BasePhpParserListener) EnterAssignmentList(ctx *AssignmentListContext) {}

// ExitAssignmentList is called when production assignmentList is exited.
func (s *BasePhpParserListener) ExitAssignmentList(ctx *AssignmentListContext) {}

// EnterAssignmentListElement is called when production assignmentListElement is entered.
func (s *BasePhpParserListener) EnterAssignmentListElement(ctx *AssignmentListElementContext) {}

// ExitAssignmentListElement is called when production assignmentListElement is exited.
func (s *BasePhpParserListener) ExitAssignmentListElement(ctx *AssignmentListElementContext) {}

// EnterModifier is called when production modifier is entered.
func (s *BasePhpParserListener) EnterModifier(ctx *ModifierContext) {}

// ExitModifier is called when production modifier is exited.
func (s *BasePhpParserListener) ExitModifier(ctx *ModifierContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BasePhpParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BasePhpParserListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterMemberModifier is called when production memberModifier is entered.
func (s *BasePhpParserListener) EnterMemberModifier(ctx *MemberModifierContext) {}

// ExitMemberModifier is called when production memberModifier is exited.
func (s *BasePhpParserListener) ExitMemberModifier(ctx *MemberModifierContext) {}

// EnterMagicConstant is called when production magicConstant is entered.
func (s *BasePhpParserListener) EnterMagicConstant(ctx *MagicConstantContext) {}

// ExitMagicConstant is called when production magicConstant is exited.
func (s *BasePhpParserListener) ExitMagicConstant(ctx *MagicConstantContext) {}

// EnterMagicMethod is called when production magicMethod is entered.
func (s *BasePhpParserListener) EnterMagicMethod(ctx *MagicMethodContext) {}

// ExitMagicMethod is called when production magicMethod is exited.
func (s *BasePhpParserListener) ExitMagicMethod(ctx *MagicMethodContext) {}

// EnterPrimitiveType is called when production primitiveType is entered.
func (s *BasePhpParserListener) EnterPrimitiveType(ctx *PrimitiveTypeContext) {}

// ExitPrimitiveType is called when production primitiveType is exited.
func (s *BasePhpParserListener) ExitPrimitiveType(ctx *PrimitiveTypeContext) {}

// EnterCastOperation is called when production castOperation is entered.
func (s *BasePhpParserListener) EnterCastOperation(ctx *CastOperationContext) {}

// ExitCastOperation is called when production castOperation is exited.
func (s *BasePhpParserListener) ExitCastOperation(ctx *CastOperationContext) {}
