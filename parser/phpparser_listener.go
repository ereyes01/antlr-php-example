// Code generated from PhpParser.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // PhpParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PhpParserListener is a complete listener for a parse tree produced by PhpParser.
type PhpParserListener interface {
	antlr.ParseTreeListener

	// EnterHtmlDocument is called when entering the htmlDocument production.
	EnterHtmlDocument(c *HtmlDocumentContext)

	// EnterHtmlElementOrPhpBlock is called when entering the htmlElementOrPhpBlock production.
	EnterHtmlElementOrPhpBlock(c *HtmlElementOrPhpBlockContext)

	// EnterHtmlElements is called when entering the htmlElements production.
	EnterHtmlElements(c *HtmlElementsContext)

	// EnterHtmlElement is called when entering the htmlElement production.
	EnterHtmlElement(c *HtmlElementContext)

	// EnterScriptTextPart is called when entering the scriptTextPart production.
	EnterScriptTextPart(c *ScriptTextPartContext)

	// EnterPhpBlock is called when entering the phpBlock production.
	EnterPhpBlock(c *PhpBlockContext)

	// EnterImportStatement is called when entering the importStatement production.
	EnterImportStatement(c *ImportStatementContext)

	// EnterTopStatement is called when entering the topStatement production.
	EnterTopStatement(c *TopStatementContext)

	// EnterUseDeclaration is called when entering the useDeclaration production.
	EnterUseDeclaration(c *UseDeclarationContext)

	// EnterUseDeclarationContentList is called when entering the useDeclarationContentList production.
	EnterUseDeclarationContentList(c *UseDeclarationContentListContext)

	// EnterUseDeclarationContent is called when entering the useDeclarationContent production.
	EnterUseDeclarationContent(c *UseDeclarationContentContext)

	// EnterNamespaceDeclaration is called when entering the namespaceDeclaration production.
	EnterNamespaceDeclaration(c *NamespaceDeclarationContext)

	// EnterNamespaceStatement is called when entering the namespaceStatement production.
	EnterNamespaceStatement(c *NamespaceStatementContext)

	// EnterFunctionDeclaration is called when entering the functionDeclaration production.
	EnterFunctionDeclaration(c *FunctionDeclarationContext)

	// EnterClassDeclaration is called when entering the classDeclaration production.
	EnterClassDeclaration(c *ClassDeclarationContext)

	// EnterClassEntryType is called when entering the classEntryType production.
	EnterClassEntryType(c *ClassEntryTypeContext)

	// EnterInterfaceList is called when entering the interfaceList production.
	EnterInterfaceList(c *InterfaceListContext)

	// EnterTypeParameterListInBrackets is called when entering the typeParameterListInBrackets production.
	EnterTypeParameterListInBrackets(c *TypeParameterListInBracketsContext)

	// EnterTypeParameterList is called when entering the typeParameterList production.
	EnterTypeParameterList(c *TypeParameterListContext)

	// EnterTypeParameterWithDefaultsList is called when entering the typeParameterWithDefaultsList production.
	EnterTypeParameterWithDefaultsList(c *TypeParameterWithDefaultsListContext)

	// EnterTypeParameterDecl is called when entering the typeParameterDecl production.
	EnterTypeParameterDecl(c *TypeParameterDeclContext)

	// EnterTypeParameterWithDefaultDecl is called when entering the typeParameterWithDefaultDecl production.
	EnterTypeParameterWithDefaultDecl(c *TypeParameterWithDefaultDeclContext)

	// EnterGenericDynamicArgs is called when entering the genericDynamicArgs production.
	EnterGenericDynamicArgs(c *GenericDynamicArgsContext)

	// EnterAttributes is called when entering the attributes production.
	EnterAttributes(c *AttributesContext)

	// EnterAttributesGroup is called when entering the attributesGroup production.
	EnterAttributesGroup(c *AttributesGroupContext)

	// EnterAttribute is called when entering the attribute production.
	EnterAttribute(c *AttributeContext)

	// EnterAttributeArgList is called when entering the attributeArgList production.
	EnterAttributeArgList(c *AttributeArgListContext)

	// EnterAttributeNamedArgList is called when entering the attributeNamedArgList production.
	EnterAttributeNamedArgList(c *AttributeNamedArgListContext)

	// EnterAttributeNamedArg is called when entering the attributeNamedArg production.
	EnterAttributeNamedArg(c *AttributeNamedArgContext)

	// EnterInnerStatementList is called when entering the innerStatementList production.
	EnterInnerStatementList(c *InnerStatementListContext)

	// EnterInnerStatement is called when entering the innerStatement production.
	EnterInnerStatement(c *InnerStatementContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterVoidStatement is called when entering the voidStatement production.
	EnterVoidStatement(c *VoidStatementContext)

	// EnterBlockStatement is called when entering the blockStatement production.
	EnterBlockStatement(c *BlockStatementContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterElseIfStatement is called when entering the elseIfStatement production.
	EnterElseIfStatement(c *ElseIfStatementContext)

	// EnterElseIfColonStatement is called when entering the elseIfColonStatement production.
	EnterElseIfColonStatement(c *ElseIfColonStatementContext)

	// EnterElseStatement is called when entering the elseStatement production.
	EnterElseStatement(c *ElseStatementContext)

	// EnterElseColonStatement is called when entering the elseColonStatement production.
	EnterElseColonStatement(c *ElseColonStatementContext)

	// EnterWhileStatement is called when entering the whileStatement production.
	EnterWhileStatement(c *WhileStatementContext)

	// EnterDoWhileStatement is called when entering the doWhileStatement production.
	EnterDoWhileStatement(c *DoWhileStatementContext)

	// EnterForStatement is called when entering the forStatement production.
	EnterForStatement(c *ForStatementContext)

	// EnterForInit is called when entering the forInit production.
	EnterForInit(c *ForInitContext)

	// EnterForUpdate is called when entering the forUpdate production.
	EnterForUpdate(c *ForUpdateContext)

	// EnterSwitchStatement is called when entering the switchStatement production.
	EnterSwitchStatement(c *SwitchStatementContext)

	// EnterSwitchBlock is called when entering the switchBlock production.
	EnterSwitchBlock(c *SwitchBlockContext)

	// EnterBreakStatement is called when entering the breakStatement production.
	EnterBreakStatement(c *BreakStatementContext)

	// EnterContinueStatement is called when entering the continueStatement production.
	EnterContinueStatement(c *ContinueStatementContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterExpressionStatement is called when entering the expressionStatement production.
	EnterExpressionStatement(c *ExpressionStatementContext)

	// EnterUnsetStatement is called when entering the unsetStatement production.
	EnterUnsetStatement(c *UnsetStatementContext)

	// EnterForeachStatement is called when entering the foreachStatement production.
	EnterForeachStatement(c *ForeachStatementContext)

	// EnterTryCatchFinally is called when entering the tryCatchFinally production.
	EnterTryCatchFinally(c *TryCatchFinallyContext)

	// EnterCatchClause is called when entering the catchClause production.
	EnterCatchClause(c *CatchClauseContext)

	// EnterFinallyStatement is called when entering the finallyStatement production.
	EnterFinallyStatement(c *FinallyStatementContext)

	// EnterThrowStatement is called when entering the throwStatement production.
	EnterThrowStatement(c *ThrowStatementContext)

	// EnterGotoStatement is called when entering the gotoStatement production.
	EnterGotoStatement(c *GotoStatementContext)

	// EnterDeclareStatement is called when entering the declareStatement production.
	EnterDeclareStatement(c *DeclareStatementContext)

	// EnterInlineHtmlStatement is called when entering the inlineHtmlStatement production.
	EnterInlineHtmlStatement(c *InlineHtmlStatementContext)

	// EnterInlineHtml is called when entering the inlineHtml production.
	EnterInlineHtml(c *InlineHtmlContext)

	// EnterDeclareList is called when entering the declareList production.
	EnterDeclareList(c *DeclareListContext)

	// EnterFormalParameterList is called when entering the formalParameterList production.
	EnterFormalParameterList(c *FormalParameterListContext)

	// EnterFormalParameter is called when entering the formalParameter production.
	EnterFormalParameter(c *FormalParameterContext)

	// EnterTypeHint is called when entering the typeHint production.
	EnterTypeHint(c *TypeHintContext)

	// EnterGlobalStatement is called when entering the globalStatement production.
	EnterGlobalStatement(c *GlobalStatementContext)

	// EnterGlobalVar is called when entering the globalVar production.
	EnterGlobalVar(c *GlobalVarContext)

	// EnterEchoStatement is called when entering the echoStatement production.
	EnterEchoStatement(c *EchoStatementContext)

	// EnterStaticVariableStatement is called when entering the staticVariableStatement production.
	EnterStaticVariableStatement(c *StaticVariableStatementContext)

	// EnterClassStatement is called when entering the classStatement production.
	EnterClassStatement(c *ClassStatementContext)

	// EnterTraitAdaptations is called when entering the traitAdaptations production.
	EnterTraitAdaptations(c *TraitAdaptationsContext)

	// EnterTraitAdaptationStatement is called when entering the traitAdaptationStatement production.
	EnterTraitAdaptationStatement(c *TraitAdaptationStatementContext)

	// EnterTraitPrecedence is called when entering the traitPrecedence production.
	EnterTraitPrecedence(c *TraitPrecedenceContext)

	// EnterTraitAlias is called when entering the traitAlias production.
	EnterTraitAlias(c *TraitAliasContext)

	// EnterTraitMethodReference is called when entering the traitMethodReference production.
	EnterTraitMethodReference(c *TraitMethodReferenceContext)

	// EnterBaseCtorCall is called when entering the baseCtorCall production.
	EnterBaseCtorCall(c *BaseCtorCallContext)

	// EnterMethodBody is called when entering the methodBody production.
	EnterMethodBody(c *MethodBodyContext)

	// EnterPropertyModifiers is called when entering the propertyModifiers production.
	EnterPropertyModifiers(c *PropertyModifiersContext)

	// EnterMemberModifiers is called when entering the memberModifiers production.
	EnterMemberModifiers(c *MemberModifiersContext)

	// EnterVariableInitializer is called when entering the variableInitializer production.
	EnterVariableInitializer(c *VariableInitializerContext)

	// EnterIdentifierInititalizer is called when entering the identifierInititalizer production.
	EnterIdentifierInititalizer(c *IdentifierInititalizerContext)

	// EnterGlobalConstantDeclaration is called when entering the globalConstantDeclaration production.
	EnterGlobalConstantDeclaration(c *GlobalConstantDeclarationContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterParenthesis is called when entering the parenthesis production.
	EnterParenthesis(c *ParenthesisContext)

	// EnterChainExpression is called when entering the ChainExpression production.
	EnterChainExpression(c *ChainExpressionContext)

	// EnterUnaryOperatorExpression is called when entering the UnaryOperatorExpression production.
	EnterUnaryOperatorExpression(c *UnaryOperatorExpressionContext)

	// EnterSpecialWordExpression is called when entering the SpecialWordExpression production.
	EnterSpecialWordExpression(c *SpecialWordExpressionContext)

	// EnterArrayCreationExpression is called when entering the ArrayCreationExpression production.
	EnterArrayCreationExpression(c *ArrayCreationExpressionContext)

	// EnterParenthesisExpression is called when entering the ParenthesisExpression production.
	EnterParenthesisExpression(c *ParenthesisExpressionContext)

	// EnterBackQuoteStringExpression is called when entering the BackQuoteStringExpression production.
	EnterBackQuoteStringExpression(c *BackQuoteStringExpressionContext)

	// EnterConditionalExpression is called when entering the ConditionalExpression production.
	EnterConditionalExpression(c *ConditionalExpressionContext)

	// EnterNewExprClause is called when entering the NewExprClause production.
	EnterNewExprClause(c *NewExprClauseContext)

	// EnterArithmeticExpression is called when entering the ArithmeticExpression production.
	EnterArithmeticExpression(c *ArithmeticExpressionContext)

	// EnterIndexerExpression is called when entering the IndexerExpression production.
	EnterIndexerExpression(c *IndexerExpressionContext)

	// EnterScalarExpression is called when entering the ScalarExpression production.
	EnterScalarExpression(c *ScalarExpressionContext)

	// EnterPrefixIncDecExpression is called when entering the PrefixIncDecExpression production.
	EnterPrefixIncDecExpression(c *PrefixIncDecExpressionContext)

	// EnterComparisonExpression is called when entering the ComparisonExpression production.
	EnterComparisonExpression(c *ComparisonExpressionContext)

	// EnterLogicalExpression is called when entering the LogicalExpression production.
	EnterLogicalExpression(c *LogicalExpressionContext)

	// EnterPrintExpression is called when entering the PrintExpression production.
	EnterPrintExpression(c *PrintExpressionContext)

	// EnterAssignmentExpression is called when entering the AssignmentExpression production.
	EnterAssignmentExpression(c *AssignmentExpressionContext)

	// EnterPostfixIncDecExpression is called when entering the PostfixIncDecExpression production.
	EnterPostfixIncDecExpression(c *PostfixIncDecExpressionContext)

	// EnterCastExpression is called when entering the CastExpression production.
	EnterCastExpression(c *CastExpressionContext)

	// EnterInstanceOfExpression is called when entering the InstanceOfExpression production.
	EnterInstanceOfExpression(c *InstanceOfExpressionContext)

	// EnterLambdaFunctionExpression is called when entering the LambdaFunctionExpression production.
	EnterLambdaFunctionExpression(c *LambdaFunctionExpressionContext)

	// EnterBitwiseExpression is called when entering the BitwiseExpression production.
	EnterBitwiseExpression(c *BitwiseExpressionContext)

	// EnterCloneExpression is called when entering the CloneExpression production.
	EnterCloneExpression(c *CloneExpressionContext)

	// EnterNewExpr is called when entering the newExpr production.
	EnterNewExpr(c *NewExprContext)

	// EnterAssignmentOperator is called when entering the assignmentOperator production.
	EnterAssignmentOperator(c *AssignmentOperatorContext)

	// EnterYieldExpression is called when entering the yieldExpression production.
	EnterYieldExpression(c *YieldExpressionContext)

	// EnterArrayItemList is called when entering the arrayItemList production.
	EnterArrayItemList(c *ArrayItemListContext)

	// EnterArrayItem is called when entering the arrayItem production.
	EnterArrayItem(c *ArrayItemContext)

	// EnterLambdaFunctionUseVars is called when entering the lambdaFunctionUseVars production.
	EnterLambdaFunctionUseVars(c *LambdaFunctionUseVarsContext)

	// EnterLambdaFunctionUseVar is called when entering the lambdaFunctionUseVar production.
	EnterLambdaFunctionUseVar(c *LambdaFunctionUseVarContext)

	// EnterQualifiedStaticTypeRef is called when entering the qualifiedStaticTypeRef production.
	EnterQualifiedStaticTypeRef(c *QualifiedStaticTypeRefContext)

	// EnterTypeRef is called when entering the typeRef production.
	EnterTypeRef(c *TypeRefContext)

	// EnterIndirectTypeRef is called when entering the indirectTypeRef production.
	EnterIndirectTypeRef(c *IndirectTypeRefContext)

	// EnterQualifiedNamespaceName is called when entering the qualifiedNamespaceName production.
	EnterQualifiedNamespaceName(c *QualifiedNamespaceNameContext)

	// EnterNamespaceNameList is called when entering the namespaceNameList production.
	EnterNamespaceNameList(c *NamespaceNameListContext)

	// EnterQualifiedNamespaceNameList is called when entering the qualifiedNamespaceNameList production.
	EnterQualifiedNamespaceNameList(c *QualifiedNamespaceNameListContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterActualArgument is called when entering the actualArgument production.
	EnterActualArgument(c *ActualArgumentContext)

	// EnterConstantInititalizer is called when entering the constantInititalizer production.
	EnterConstantInititalizer(c *ConstantInititalizerContext)

	// EnterConstantArrayItemList is called when entering the constantArrayItemList production.
	EnterConstantArrayItemList(c *ConstantArrayItemListContext)

	// EnterConstantArrayItem is called when entering the constantArrayItem production.
	EnterConstantArrayItem(c *ConstantArrayItemContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterLiteralConstant is called when entering the literalConstant production.
	EnterLiteralConstant(c *LiteralConstantContext)

	// EnterNumericConstant is called when entering the numericConstant production.
	EnterNumericConstant(c *NumericConstantContext)

	// EnterClassConstant is called when entering the classConstant production.
	EnterClassConstant(c *ClassConstantContext)

	// EnterStringConstant is called when entering the stringConstant production.
	EnterStringConstant(c *StringConstantContext)

	// EnterStr is called when entering the str production.
	EnterStr(c *StrContext)

	// EnterInterpolatedStringPart is called when entering the interpolatedStringPart production.
	EnterInterpolatedStringPart(c *InterpolatedStringPartContext)

	// EnterChainList is called when entering the chainList production.
	EnterChainList(c *ChainListContext)

	// EnterChain is called when entering the chain production.
	EnterChain(c *ChainContext)

	// EnterMemberAccess is called when entering the memberAccess production.
	EnterMemberAccess(c *MemberAccessContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterFunctionCallName is called when entering the functionCallName production.
	EnterFunctionCallName(c *FunctionCallNameContext)

	// EnterActualArguments is called when entering the actualArguments production.
	EnterActualArguments(c *ActualArgumentsContext)

	// EnterChainBase is called when entering the chainBase production.
	EnterChainBase(c *ChainBaseContext)

	// EnterKeyedFieldName is called when entering the keyedFieldName production.
	EnterKeyedFieldName(c *KeyedFieldNameContext)

	// EnterKeyedSimpleFieldName is called when entering the keyedSimpleFieldName production.
	EnterKeyedSimpleFieldName(c *KeyedSimpleFieldNameContext)

	// EnterKeyedVariable is called when entering the keyedVariable production.
	EnterKeyedVariable(c *KeyedVariableContext)

	// EnterSquareCurlyExpression is called when entering the squareCurlyExpression production.
	EnterSquareCurlyExpression(c *SquareCurlyExpressionContext)

	// EnterAssignmentList is called when entering the assignmentList production.
	EnterAssignmentList(c *AssignmentListContext)

	// EnterAssignmentListElement is called when entering the assignmentListElement production.
	EnterAssignmentListElement(c *AssignmentListElementContext)

	// EnterModifier is called when entering the modifier production.
	EnterModifier(c *ModifierContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterMemberModifier is called when entering the memberModifier production.
	EnterMemberModifier(c *MemberModifierContext)

	// EnterMagicConstant is called when entering the magicConstant production.
	EnterMagicConstant(c *MagicConstantContext)

	// EnterMagicMethod is called when entering the magicMethod production.
	EnterMagicMethod(c *MagicMethodContext)

	// EnterPrimitiveType is called when entering the primitiveType production.
	EnterPrimitiveType(c *PrimitiveTypeContext)

	// EnterCastOperation is called when entering the castOperation production.
	EnterCastOperation(c *CastOperationContext)

	// ExitHtmlDocument is called when exiting the htmlDocument production.
	ExitHtmlDocument(c *HtmlDocumentContext)

	// ExitHtmlElementOrPhpBlock is called when exiting the htmlElementOrPhpBlock production.
	ExitHtmlElementOrPhpBlock(c *HtmlElementOrPhpBlockContext)

	// ExitHtmlElements is called when exiting the htmlElements production.
	ExitHtmlElements(c *HtmlElementsContext)

	// ExitHtmlElement is called when exiting the htmlElement production.
	ExitHtmlElement(c *HtmlElementContext)

	// ExitScriptTextPart is called when exiting the scriptTextPart production.
	ExitScriptTextPart(c *ScriptTextPartContext)

	// ExitPhpBlock is called when exiting the phpBlock production.
	ExitPhpBlock(c *PhpBlockContext)

	// ExitImportStatement is called when exiting the importStatement production.
	ExitImportStatement(c *ImportStatementContext)

	// ExitTopStatement is called when exiting the topStatement production.
	ExitTopStatement(c *TopStatementContext)

	// ExitUseDeclaration is called when exiting the useDeclaration production.
	ExitUseDeclaration(c *UseDeclarationContext)

	// ExitUseDeclarationContentList is called when exiting the useDeclarationContentList production.
	ExitUseDeclarationContentList(c *UseDeclarationContentListContext)

	// ExitUseDeclarationContent is called when exiting the useDeclarationContent production.
	ExitUseDeclarationContent(c *UseDeclarationContentContext)

	// ExitNamespaceDeclaration is called when exiting the namespaceDeclaration production.
	ExitNamespaceDeclaration(c *NamespaceDeclarationContext)

	// ExitNamespaceStatement is called when exiting the namespaceStatement production.
	ExitNamespaceStatement(c *NamespaceStatementContext)

	// ExitFunctionDeclaration is called when exiting the functionDeclaration production.
	ExitFunctionDeclaration(c *FunctionDeclarationContext)

	// ExitClassDeclaration is called when exiting the classDeclaration production.
	ExitClassDeclaration(c *ClassDeclarationContext)

	// ExitClassEntryType is called when exiting the classEntryType production.
	ExitClassEntryType(c *ClassEntryTypeContext)

	// ExitInterfaceList is called when exiting the interfaceList production.
	ExitInterfaceList(c *InterfaceListContext)

	// ExitTypeParameterListInBrackets is called when exiting the typeParameterListInBrackets production.
	ExitTypeParameterListInBrackets(c *TypeParameterListInBracketsContext)

	// ExitTypeParameterList is called when exiting the typeParameterList production.
	ExitTypeParameterList(c *TypeParameterListContext)

	// ExitTypeParameterWithDefaultsList is called when exiting the typeParameterWithDefaultsList production.
	ExitTypeParameterWithDefaultsList(c *TypeParameterWithDefaultsListContext)

	// ExitTypeParameterDecl is called when exiting the typeParameterDecl production.
	ExitTypeParameterDecl(c *TypeParameterDeclContext)

	// ExitTypeParameterWithDefaultDecl is called when exiting the typeParameterWithDefaultDecl production.
	ExitTypeParameterWithDefaultDecl(c *TypeParameterWithDefaultDeclContext)

	// ExitGenericDynamicArgs is called when exiting the genericDynamicArgs production.
	ExitGenericDynamicArgs(c *GenericDynamicArgsContext)

	// ExitAttributes is called when exiting the attributes production.
	ExitAttributes(c *AttributesContext)

	// ExitAttributesGroup is called when exiting the attributesGroup production.
	ExitAttributesGroup(c *AttributesGroupContext)

	// ExitAttribute is called when exiting the attribute production.
	ExitAttribute(c *AttributeContext)

	// ExitAttributeArgList is called when exiting the attributeArgList production.
	ExitAttributeArgList(c *AttributeArgListContext)

	// ExitAttributeNamedArgList is called when exiting the attributeNamedArgList production.
	ExitAttributeNamedArgList(c *AttributeNamedArgListContext)

	// ExitAttributeNamedArg is called when exiting the attributeNamedArg production.
	ExitAttributeNamedArg(c *AttributeNamedArgContext)

	// ExitInnerStatementList is called when exiting the innerStatementList production.
	ExitInnerStatementList(c *InnerStatementListContext)

	// ExitInnerStatement is called when exiting the innerStatement production.
	ExitInnerStatement(c *InnerStatementContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitVoidStatement is called when exiting the voidStatement production.
	ExitVoidStatement(c *VoidStatementContext)

	// ExitBlockStatement is called when exiting the blockStatement production.
	ExitBlockStatement(c *BlockStatementContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitElseIfStatement is called when exiting the elseIfStatement production.
	ExitElseIfStatement(c *ElseIfStatementContext)

	// ExitElseIfColonStatement is called when exiting the elseIfColonStatement production.
	ExitElseIfColonStatement(c *ElseIfColonStatementContext)

	// ExitElseStatement is called when exiting the elseStatement production.
	ExitElseStatement(c *ElseStatementContext)

	// ExitElseColonStatement is called when exiting the elseColonStatement production.
	ExitElseColonStatement(c *ElseColonStatementContext)

	// ExitWhileStatement is called when exiting the whileStatement production.
	ExitWhileStatement(c *WhileStatementContext)

	// ExitDoWhileStatement is called when exiting the doWhileStatement production.
	ExitDoWhileStatement(c *DoWhileStatementContext)

	// ExitForStatement is called when exiting the forStatement production.
	ExitForStatement(c *ForStatementContext)

	// ExitForInit is called when exiting the forInit production.
	ExitForInit(c *ForInitContext)

	// ExitForUpdate is called when exiting the forUpdate production.
	ExitForUpdate(c *ForUpdateContext)

	// ExitSwitchStatement is called when exiting the switchStatement production.
	ExitSwitchStatement(c *SwitchStatementContext)

	// ExitSwitchBlock is called when exiting the switchBlock production.
	ExitSwitchBlock(c *SwitchBlockContext)

	// ExitBreakStatement is called when exiting the breakStatement production.
	ExitBreakStatement(c *BreakStatementContext)

	// ExitContinueStatement is called when exiting the continueStatement production.
	ExitContinueStatement(c *ContinueStatementContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitExpressionStatement is called when exiting the expressionStatement production.
	ExitExpressionStatement(c *ExpressionStatementContext)

	// ExitUnsetStatement is called when exiting the unsetStatement production.
	ExitUnsetStatement(c *UnsetStatementContext)

	// ExitForeachStatement is called when exiting the foreachStatement production.
	ExitForeachStatement(c *ForeachStatementContext)

	// ExitTryCatchFinally is called when exiting the tryCatchFinally production.
	ExitTryCatchFinally(c *TryCatchFinallyContext)

	// ExitCatchClause is called when exiting the catchClause production.
	ExitCatchClause(c *CatchClauseContext)

	// ExitFinallyStatement is called when exiting the finallyStatement production.
	ExitFinallyStatement(c *FinallyStatementContext)

	// ExitThrowStatement is called when exiting the throwStatement production.
	ExitThrowStatement(c *ThrowStatementContext)

	// ExitGotoStatement is called when exiting the gotoStatement production.
	ExitGotoStatement(c *GotoStatementContext)

	// ExitDeclareStatement is called when exiting the declareStatement production.
	ExitDeclareStatement(c *DeclareStatementContext)

	// ExitInlineHtmlStatement is called when exiting the inlineHtmlStatement production.
	ExitInlineHtmlStatement(c *InlineHtmlStatementContext)

	// ExitInlineHtml is called when exiting the inlineHtml production.
	ExitInlineHtml(c *InlineHtmlContext)

	// ExitDeclareList is called when exiting the declareList production.
	ExitDeclareList(c *DeclareListContext)

	// ExitFormalParameterList is called when exiting the formalParameterList production.
	ExitFormalParameterList(c *FormalParameterListContext)

	// ExitFormalParameter is called when exiting the formalParameter production.
	ExitFormalParameter(c *FormalParameterContext)

	// ExitTypeHint is called when exiting the typeHint production.
	ExitTypeHint(c *TypeHintContext)

	// ExitGlobalStatement is called when exiting the globalStatement production.
	ExitGlobalStatement(c *GlobalStatementContext)

	// ExitGlobalVar is called when exiting the globalVar production.
	ExitGlobalVar(c *GlobalVarContext)

	// ExitEchoStatement is called when exiting the echoStatement production.
	ExitEchoStatement(c *EchoStatementContext)

	// ExitStaticVariableStatement is called when exiting the staticVariableStatement production.
	ExitStaticVariableStatement(c *StaticVariableStatementContext)

	// ExitClassStatement is called when exiting the classStatement production.
	ExitClassStatement(c *ClassStatementContext)

	// ExitTraitAdaptations is called when exiting the traitAdaptations production.
	ExitTraitAdaptations(c *TraitAdaptationsContext)

	// ExitTraitAdaptationStatement is called when exiting the traitAdaptationStatement production.
	ExitTraitAdaptationStatement(c *TraitAdaptationStatementContext)

	// ExitTraitPrecedence is called when exiting the traitPrecedence production.
	ExitTraitPrecedence(c *TraitPrecedenceContext)

	// ExitTraitAlias is called when exiting the traitAlias production.
	ExitTraitAlias(c *TraitAliasContext)

	// ExitTraitMethodReference is called when exiting the traitMethodReference production.
	ExitTraitMethodReference(c *TraitMethodReferenceContext)

	// ExitBaseCtorCall is called when exiting the baseCtorCall production.
	ExitBaseCtorCall(c *BaseCtorCallContext)

	// ExitMethodBody is called when exiting the methodBody production.
	ExitMethodBody(c *MethodBodyContext)

	// ExitPropertyModifiers is called when exiting the propertyModifiers production.
	ExitPropertyModifiers(c *PropertyModifiersContext)

	// ExitMemberModifiers is called when exiting the memberModifiers production.
	ExitMemberModifiers(c *MemberModifiersContext)

	// ExitVariableInitializer is called when exiting the variableInitializer production.
	ExitVariableInitializer(c *VariableInitializerContext)

	// ExitIdentifierInititalizer is called when exiting the identifierInititalizer production.
	ExitIdentifierInititalizer(c *IdentifierInititalizerContext)

	// ExitGlobalConstantDeclaration is called when exiting the globalConstantDeclaration production.
	ExitGlobalConstantDeclaration(c *GlobalConstantDeclarationContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitParenthesis is called when exiting the parenthesis production.
	ExitParenthesis(c *ParenthesisContext)

	// ExitChainExpression is called when exiting the ChainExpression production.
	ExitChainExpression(c *ChainExpressionContext)

	// ExitUnaryOperatorExpression is called when exiting the UnaryOperatorExpression production.
	ExitUnaryOperatorExpression(c *UnaryOperatorExpressionContext)

	// ExitSpecialWordExpression is called when exiting the SpecialWordExpression production.
	ExitSpecialWordExpression(c *SpecialWordExpressionContext)

	// ExitArrayCreationExpression is called when exiting the ArrayCreationExpression production.
	ExitArrayCreationExpression(c *ArrayCreationExpressionContext)

	// ExitParenthesisExpression is called when exiting the ParenthesisExpression production.
	ExitParenthesisExpression(c *ParenthesisExpressionContext)

	// ExitBackQuoteStringExpression is called when exiting the BackQuoteStringExpression production.
	ExitBackQuoteStringExpression(c *BackQuoteStringExpressionContext)

	// ExitConditionalExpression is called when exiting the ConditionalExpression production.
	ExitConditionalExpression(c *ConditionalExpressionContext)

	// ExitNewExprClause is called when exiting the NewExprClause production.
	ExitNewExprClause(c *NewExprClauseContext)

	// ExitArithmeticExpression is called when exiting the ArithmeticExpression production.
	ExitArithmeticExpression(c *ArithmeticExpressionContext)

	// ExitIndexerExpression is called when exiting the IndexerExpression production.
	ExitIndexerExpression(c *IndexerExpressionContext)

	// ExitScalarExpression is called when exiting the ScalarExpression production.
	ExitScalarExpression(c *ScalarExpressionContext)

	// ExitPrefixIncDecExpression is called when exiting the PrefixIncDecExpression production.
	ExitPrefixIncDecExpression(c *PrefixIncDecExpressionContext)

	// ExitComparisonExpression is called when exiting the ComparisonExpression production.
	ExitComparisonExpression(c *ComparisonExpressionContext)

	// ExitLogicalExpression is called when exiting the LogicalExpression production.
	ExitLogicalExpression(c *LogicalExpressionContext)

	// ExitPrintExpression is called when exiting the PrintExpression production.
	ExitPrintExpression(c *PrintExpressionContext)

	// ExitAssignmentExpression is called when exiting the AssignmentExpression production.
	ExitAssignmentExpression(c *AssignmentExpressionContext)

	// ExitPostfixIncDecExpression is called when exiting the PostfixIncDecExpression production.
	ExitPostfixIncDecExpression(c *PostfixIncDecExpressionContext)

	// ExitCastExpression is called when exiting the CastExpression production.
	ExitCastExpression(c *CastExpressionContext)

	// ExitInstanceOfExpression is called when exiting the InstanceOfExpression production.
	ExitInstanceOfExpression(c *InstanceOfExpressionContext)

	// ExitLambdaFunctionExpression is called when exiting the LambdaFunctionExpression production.
	ExitLambdaFunctionExpression(c *LambdaFunctionExpressionContext)

	// ExitBitwiseExpression is called when exiting the BitwiseExpression production.
	ExitBitwiseExpression(c *BitwiseExpressionContext)

	// ExitCloneExpression is called when exiting the CloneExpression production.
	ExitCloneExpression(c *CloneExpressionContext)

	// ExitNewExpr is called when exiting the newExpr production.
	ExitNewExpr(c *NewExprContext)

	// ExitAssignmentOperator is called when exiting the assignmentOperator production.
	ExitAssignmentOperator(c *AssignmentOperatorContext)

	// ExitYieldExpression is called when exiting the yieldExpression production.
	ExitYieldExpression(c *YieldExpressionContext)

	// ExitArrayItemList is called when exiting the arrayItemList production.
	ExitArrayItemList(c *ArrayItemListContext)

	// ExitArrayItem is called when exiting the arrayItem production.
	ExitArrayItem(c *ArrayItemContext)

	// ExitLambdaFunctionUseVars is called when exiting the lambdaFunctionUseVars production.
	ExitLambdaFunctionUseVars(c *LambdaFunctionUseVarsContext)

	// ExitLambdaFunctionUseVar is called when exiting the lambdaFunctionUseVar production.
	ExitLambdaFunctionUseVar(c *LambdaFunctionUseVarContext)

	// ExitQualifiedStaticTypeRef is called when exiting the qualifiedStaticTypeRef production.
	ExitQualifiedStaticTypeRef(c *QualifiedStaticTypeRefContext)

	// ExitTypeRef is called when exiting the typeRef production.
	ExitTypeRef(c *TypeRefContext)

	// ExitIndirectTypeRef is called when exiting the indirectTypeRef production.
	ExitIndirectTypeRef(c *IndirectTypeRefContext)

	// ExitQualifiedNamespaceName is called when exiting the qualifiedNamespaceName production.
	ExitQualifiedNamespaceName(c *QualifiedNamespaceNameContext)

	// ExitNamespaceNameList is called when exiting the namespaceNameList production.
	ExitNamespaceNameList(c *NamespaceNameListContext)

	// ExitQualifiedNamespaceNameList is called when exiting the qualifiedNamespaceNameList production.
	ExitQualifiedNamespaceNameList(c *QualifiedNamespaceNameListContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitActualArgument is called when exiting the actualArgument production.
	ExitActualArgument(c *ActualArgumentContext)

	// ExitConstantInititalizer is called when exiting the constantInititalizer production.
	ExitConstantInititalizer(c *ConstantInititalizerContext)

	// ExitConstantArrayItemList is called when exiting the constantArrayItemList production.
	ExitConstantArrayItemList(c *ConstantArrayItemListContext)

	// ExitConstantArrayItem is called when exiting the constantArrayItem production.
	ExitConstantArrayItem(c *ConstantArrayItemContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitLiteralConstant is called when exiting the literalConstant production.
	ExitLiteralConstant(c *LiteralConstantContext)

	// ExitNumericConstant is called when exiting the numericConstant production.
	ExitNumericConstant(c *NumericConstantContext)

	// ExitClassConstant is called when exiting the classConstant production.
	ExitClassConstant(c *ClassConstantContext)

	// ExitStringConstant is called when exiting the stringConstant production.
	ExitStringConstant(c *StringConstantContext)

	// ExitStr is called when exiting the str production.
	ExitStr(c *StrContext)

	// ExitInterpolatedStringPart is called when exiting the interpolatedStringPart production.
	ExitInterpolatedStringPart(c *InterpolatedStringPartContext)

	// ExitChainList is called when exiting the chainList production.
	ExitChainList(c *ChainListContext)

	// ExitChain is called when exiting the chain production.
	ExitChain(c *ChainContext)

	// ExitMemberAccess is called when exiting the memberAccess production.
	ExitMemberAccess(c *MemberAccessContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitFunctionCallName is called when exiting the functionCallName production.
	ExitFunctionCallName(c *FunctionCallNameContext)

	// ExitActualArguments is called when exiting the actualArguments production.
	ExitActualArguments(c *ActualArgumentsContext)

	// ExitChainBase is called when exiting the chainBase production.
	ExitChainBase(c *ChainBaseContext)

	// ExitKeyedFieldName is called when exiting the keyedFieldName production.
	ExitKeyedFieldName(c *KeyedFieldNameContext)

	// ExitKeyedSimpleFieldName is called when exiting the keyedSimpleFieldName production.
	ExitKeyedSimpleFieldName(c *KeyedSimpleFieldNameContext)

	// ExitKeyedVariable is called when exiting the keyedVariable production.
	ExitKeyedVariable(c *KeyedVariableContext)

	// ExitSquareCurlyExpression is called when exiting the squareCurlyExpression production.
	ExitSquareCurlyExpression(c *SquareCurlyExpressionContext)

	// ExitAssignmentList is called when exiting the assignmentList production.
	ExitAssignmentList(c *AssignmentListContext)

	// ExitAssignmentListElement is called when exiting the assignmentListElement production.
	ExitAssignmentListElement(c *AssignmentListElementContext)

	// ExitModifier is called when exiting the modifier production.
	ExitModifier(c *ModifierContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitMemberModifier is called when exiting the memberModifier production.
	ExitMemberModifier(c *MemberModifierContext)

	// ExitMagicConstant is called when exiting the magicConstant production.
	ExitMagicConstant(c *MagicConstantContext)

	// ExitMagicMethod is called when exiting the magicMethod production.
	ExitMagicMethod(c *MagicMethodContext)

	// ExitPrimitiveType is called when exiting the primitiveType production.
	ExitPrimitiveType(c *PrimitiveTypeContext)

	// ExitCastOperation is called when exiting the castOperation production.
	ExitCastOperation(c *CastOperationContext)
}
