package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string }
func genFn0() uint64 {
    var v4 string = "\n\t"
    v4 = v4
    v4 = ("hello" + "foo bar")
    v4 = v4
    var v5 int16 = -32768
    _ = v4
    _ = v5
    return 4294967295
}
func genFn1() string {
    var v6 int32 = 42
    if (!(true || (565 != 255))) {
        v7 := GenS1{F0: 2}
        _ = v7
    } else {
        var v8 float64 = 0.1
        _ = v8
    }
    _ = v6
    return "hello"
}

func main() {
    v1 := []string{"foo bar", "foo bar", "foo bar"}
    var v2 uint64 = 1
    v3 := make(map[string]string)
    v9 := v3[""]
    for fi10 := 0; fi10 < 15; fi10++ {
        v11 := GenS1{F0: 245}
        for fi12 := 0; fi12 < 5; fi12++ {
            v2 = 4294967295
            var v13 int32 = -2147483648
            _ = fi12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    var v14 uint64 = (^(v2))
    for fi15 := 0; fi15 < 8; fi15++ {
        var v16 int16 = 1000
        _ = len(v3)
        v2 -= v14
        _ = fi15
        _ = v16
    }
    v17 := GenS1{F0: uint8(v2)}
    if false {
        v2 = 18446744073709551615
        v2 = (v14 - v2)
    }
    v18 := GenS1{F0: 105}
    if (!(2147483647 != 2147483647)) {
        switch 127 {
        case 1:
            fmt.Println(v9, v14, v2)
        case -515:
            switch 4294967295 {
            case -9223372036854775808:
                for fi19 := 0; fi19 < 7; fi19++ {
                    if true {
                        v14 -= genFn0()
                    } else {
                        v14 -= genFn0()
                        v14 = genFn0()
                    }
                    switch "\n\t" {
                    case "foo bar":
                        if false {
                            delete(v3, v9)
                            fmt.Println(fi19)
                        }
                    case "a":
                        var v20 int8 = -128
                        _ = v20
                    default:
                        if len(v1) > 0 {
                            v1[0] = "foo bar"
                        }
                    }
                    v21 := GenS1{F0: 127}
                    _ = fi19
                    _ = v21
                }
            case 0:
                v22 := v3[v9]
                _ = v22
            case 42:
                var v23 string = (((v9 + ("a" + v9)) + "a") + (v9 + (((v9 + v9) + (((v9 + "hello") + v9) + ((v9 + "\n\t") + v9))) + v9)))
                _ = v23
            default:
                if false {
                    for fi24 := 0; fi24 < 4; fi24++ {
                        v9 = (v9 + v9)
                        fmt.Println(v14)
                        if ((!true) && (false || (!(true && (false || (!true)))))) {
                            v2 -= 530
                        } else {
                            for fi25 := 0; fi25 < 20; fi25++ {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                if len(v1) > 0 {
                                    v1[0] = v9
                                }
                                if (-20 == 46) {
                                    v2 = v14
                                    v2 = 2
                                } else {
                                    switch (fi25 - fi24) {
                                    case 4294967295:
                                        fmt.Println(fi24, v9)
                                    case 127:
                                        fmt.Println(v2, fi24)
                                    case 128:
                                        fmt.Println(fi24, v14, fi25)
                                    }
                                }
                                _ = len(v3)
                                v26 := GenS1{F0: 128}
                                _ = fi25
                                _ = v26
                            }
                        }
                        _ = fi24
                    }
                    switch 9223372036854775807 {
                    case -799:
                        v14 = 181
                    case -237:
                        if ((v2 | (^(v2))) < 42) {
                            v9 = "hello"
                        } else {
                            v14 = genFn0()
                            v14 = v14
                        }
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    var v27 int16 = 333
                    _ = v27
                }
            }
        case 0:
            v2 += 2
        default:
            v28 := GenS1{F0: 58}
            _ = v28
        }
    } else {
        fmt.Println(v2, v2)
    }
    v29 := GenS2{F0: v9}
    var v30 string = v9
    var v31 float32 = 1e10
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v9
}
