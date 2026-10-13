package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 string; F2 bool }
type GenS3 struct { F0 uint32 }
func genFn0(p0 int16) bool {
    p0 -= 32767
    v7 := GenS1{F0: "hello"}
    _ = p0
    _ = v7
    return false
}

func main() {
    var v1 uint8 = 255
    var v2 float32 = (3.14159 + 0.0)
    v3 := make(map[string]string)
    var v4 int64 = -9223372036854775808
    var v5 string = ("" + (("\n\t" + ("hello" + ("foo bar" + "a"))) + ((("foo bar" + "a") + "hello") + "hello")))
    var v6 int8 = -1
    if true {
        if (2147483647 == -750) {
            switch int(v6) {
            case 528:
                switch -208 {
                case -2147483648:
                    if true {
                        fmt.Println(v1, v4)
                    } else {
                        var v8 string = "hello"
                        switch "\n\t" {
                        case "hello":
                            fmt.Println(v8)
                        case "a":
                            v9 := GenS1{F0: v8}
                            _ = v9
                        default:
                            if ((true || (("\n\t" + ("\n\t" + (v8 + (v5 + "foo bar")))) <= v5)) || ((^(v1)) <= 31)) {
                                delete(v3, (v5 + ((v8 + ("" + ((v5 + v5) + "\n\t"))) + "")))
                            } else {
                                v6 = 34
                                if ((!((v4 / ((v4) | 1)) == (^(v4)))) && false) {
                                    v6 = -52
                                } else {
                                    for fi10 := 0; fi10 < 5; fi10++ {
                                        for fi11 := 0; fi11 < 2; fi11++ {
                                            v8 = v5
                                            var v12 bool = (42 == -32768)
                                            if v12 {
                                                v1 -= 1
                                            }
                                            _ = fi11
                                            _ = v12
                                        }
                                        v13 := GenS1{F0: v5}
                                        fmt.Println(v6, fi10)
                                        _ = fi10
                                        _ = v13
                                    }
                                    var v14 int32 = 436
                                    for fi15 := 0; fi15 < 16; fi15++ {
                                        var v16 int = fi15
                                        for fi17 := 0; fi17 < 9; fi17++ {
                                            v18 := GenS1{F0: (v5 + v5)}
                                            v19 := GenS1{F0: v5}
                                            _ = fi17
                                            _ = v18
                                            _ = v19
                                        }
                                        v20 := GenS1{F0: ("foo bar" + v8)}
                                        _ = fi15
                                        _ = v16
                                        _ = v20
                                    }
                                    _ = v14
                                }
                            }
                        }
                        _ = v8
                    }
                case 938:
                    var v21 int32 = 42
                    _ = v21
                case 127:
                    fmt.Println(v4)
                }
            case 4294967295:
                v2 = 0.0
            case 128:
                fmt.Println(v1)
            }
            for fi22 := 0; fi22 < 14; fi22++ {
                var v23 int64 = v4
                v23 += v23
                _ = fi22
                _ = v23
            }
            for fi24 := 0; fi24 < 20; fi24++ {
                fmt.Println(v1, v1)
                if true {
                    v2 = v2
                    var v25 int = fi24
                    v26 := GenS1{F0: ""}
                    _ = v25
                    _ = v26
                }
                var v27 bool = false
                _ = fi24
                _ = v27
            }
        }
    } else {
        switch v5 {
        case "":
            v28 := GenS2{F0: (v6 + v6), F1: (("\n\t" + v5) + ("\n\t" + ((((v5 + v5) + v5) + "") + ("\n\t" + v5)))), F2: (true || false)}
            _ = v28
        default:
            v3[("foo bar" + "a")] = v5
        }
        v29 := GenS1{F0: v5}
        _ = v29
    }
    var v30 int16 = 210
    v31 := []int{13, 127}
    for fi32 := 0; fi32 < 18; fi32++ {
        fmt.Println(fi32, v30, v1)
        _ = fi32
    }
    fmt.Println(v5, v1, v2)
    v33 := GenS1{F0: "a"}
    v1 += v1
    var v34 int8 = v6
    v6 += v34
    for fi35 := 0; fi35 < 16; fi35++ {
        for fi36 := 0; fi36 < 10; fi36++ {
            v37 := GenS3{F0: 0}
            v30 += v30
            switch fi36 {
            case 9223372036854775807:
                v38 := GenS2{F0: (^(v34)), F1: "a", F2: (0 == 1)}
                _ = v38
            case 2147483647:
                v39 := GenS3{F0: 854}
                _ = v39
            case 0:
                v40 := GenS1{F0: v5}
                _ = v40
            default:
                v5 = ("hello" + v5)
            }
            _ = fi36
            _ = v37
        }
        var v41 int8 = -10
        _ = fi35
        _ = v41
    }
    v42 := GenS2{F0: -1, F1: v5, F2: (v4 <= v4)}
    v43 := v3[v5]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v34
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
