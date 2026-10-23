package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 int32 }

func main() {
    var v1 uint16 = 1
    var v2 float64 = (1.5 - -2.25)
    var v3 float32 = 0.1
    v4 := make(map[uint32]int)
    var v5 bool = ((!(-100 > -38)) && (0 < 255))
    v6 := make(map[uint32]int)
    var v7 int8 = -128
    var v8 int32 = -744
    var v9 int8 = -128
    v7 += v9
    v8 = v8
    switch 127 {
    case 32768:
        if v5 {
            v1 += 32767
        } else {
            var v10 bool = v5
            _ = v10
        }
    default:
        switch 4294967295 {
        case -2147483648:
            v2 += 0.1
        case 127:
            v7 -= (v9 - 100)
        }
    }
    v11 := GenS1{F0: 9223372036854775807}
    delete(v4, 65535)
    switch -2147483648 {
    case 32767:
        for fi12 := 0; fi12 < 15; fi12++ {
            var v13 float64 = 3.14159
            _ = fi12
            _ = v13
        }
    }
    v14 := GenS1{F0: -9223372036854775808}
    v5 = (("" + "") >= (("a" + (("\n\t" + "\n\t") + (("\n\t" + "hello") + ("hello" + (("\n\t" + ("foo bar" + "foo bar")) + "foo bar"))))) + ("" + "foo bar")))
    switch 127 {
    case 32767:
        if (!v5) {
            v15 := v6[65535]
            v16 := v6[683]
            _ = v15
            _ = v16
        } else {
            switch -917 {
            case 0:
                for fi17 := 0; fi17 < 16; fi17++ {
                    for fi18 := 0; fi18 < 2; fi18++ {
                        v9 = (v9 & (-(v7)))
                        var v19 float32 = (v3 / (v3 + ((v3 / 1.0) + v3)))
                        if (0 == 2) {
                            switch int(v9) {
                            case -578:
                                var v20 int8 = v9
                                _ = v20
                            }
                        }
                        _ = fi18
                        _ = v19
                    }
                    var v21 uint64 = 4294967295
                    _ = fi17
                    _ = v21
                }
            case 4294967295:
                var v22 uint64 = 0
                _ = v22
            default:
                fmt.Println(v7)
            }
        }
    }
    switch "\n\t" {
    case "":
        var v23 int = 1
        _ = v23
    case "a":
        switch 32768 {
        case -9223372036854775808:
            var v24 float32 = (v3 * float32((1.5 / v2)))
            _ = v24
        }
    case "foo bar":
        if (v8 == (-(v8))) {
            for fi25 := 0; fi25 < 11; fi25++ {
                for fi26 := 0; fi26 < 6; fi26++ {
                    var v27 int64 = -610
                    _ = fi26
                    _ = v27
                }
                if true {
                    v5 = (2147483648 != uint32(v1))
                    for fi28 := 0; fi28 < 7; fi28++ {
                        fmt.Println(v2)
                        switch fi25 {
                        case 9223372036854775807:
                            v5 = (42 <= 0)
                        }
                        var v29 int16 = -32768
                        _ = fi28
                        _ = v29
                    }
                    switch 4294967295 {
                    case -2147483648:
                        switch (("hello" + ("\n\t" + ("\n\t" + ""))) + ("" + ((("foo bar" + (("\n\t" + (((((("\n\t" + "a") + ("" + "foo bar")) + "\n\t") + ("" + (("foo bar" + "foo bar") + "foo bar"))) + "\n\t") + "hello")) + "\n\t")) + "foo bar") + "hello"))) {
                        case "\n\t":
                            v30 := GenS1{F0: -757}
                            _ = v30
                        default:
                            for fi31 := 0; fi31 < 6; fi31++ {
                                var v32 uint64 = 883
                                v8 -= int32(v32)
                                var v33 uint64 = (^(v32))
                                _ = fi31
                                _ = v32
                                _ = v33
                            }
                        }
                    case -372:
                        var v34 int = fi25
                        _ = v34
                    case -1:
                        switch fi25 {
                        case 1:
                            var v35 int = fi25
                            _ = v35
                        case 2147483647:
                            v36 := []int{(^(fi25)), 2147483647, fi25, fi25}
                            _ = v36
                        }
                    }
                } else {
                    v8 -= v8
                    v2 = v2
                    switch fi25 {
                    case 32768:
                        for fi37 := 0; fi37 < 9; fi37++ {
                            fmt.Println(v2, fi25)
                            v38 := GenS1{F0: 0}
                            if ((v5 && (32767 < -1)) && v5) {
                                v39 := GenS1{F0: 9223372036854775807}
                                if (v5 && (true || v5)) {
                                    switch -152 {
                                    case 9223372036854775807:
                                        _ = len(v4)
                                    }
                                    v40 := GenS1{F0: 141}
                                    v41 := v4[42]
                                    _ = v40
                                    _ = v41
                                } else {
                                    switch (fi37 + 128) {
                                    case 0:
                                        fmt.Println(v7, v2, fi37)
                                    case 42:
                                        v42 := GenS1{F0: 4}
                                        _ = v42
                                    }
                                    fmt.Println(fi25, v2, v2)
                                    v43 := []int{fi37, (fi25 ^ (fi37 & fi25))}
                                    v44 := GenS2{F0: "a", F1: v8}
                                    _ = v43
                                    _ = v44
                                }
                                v3 += v3
                                _ = v39
                            } else {
                                v2 = (v2 - float64((v3 + v3)))
                                v6[1] = fi37
                                v7 = 42
                            }
                            _ = fi37
                            _ = v38
                        }
                    case 128:
                        for fi45 := 0; fi45 < 15; fi45++ {
                            v46 := []int{(fi25 << ((4294967295) & 63)), fi25, (fi25 * fi45)}
                            var v47 float64 = (v2 - (v2 / v2))
                            _ = fi45
                            _ = v46
                            _ = v47
                        }
                    default:
                        v48 := GenS2{F0: ("\n\t" + ("foo bar" + ("a" + "hello"))), F1: (-(v8))}
                        _ = v48
                    }
                }
                if (!false) {
                    v1 += 2
                    fmt.Println(v8)
                }
                _ = fi25
            }
            var v49 uint16 = v1
            if false {
                var v50 uint16 = 32767
                v49 = v1
                var v51 int8 = v9
                _ = v50
                _ = v51
            } else {
                fmt.Println(v5)
            }
            _ = v49
        }
    }
    fmt.Println(v1, v7)
    if (!false) {
        v52 := []int{}
        if v5 {
            v1 += v1
        } else {
            fmt.Println(v8, v5, v8)
        }
        fmt.Println(v2, v8)
        _ = v52
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
