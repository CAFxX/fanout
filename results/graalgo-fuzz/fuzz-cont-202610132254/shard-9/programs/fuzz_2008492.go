package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 uint; F1 uint; F2 uint8 }
type GenS4 struct { F0 uint; F1 int64; F2 string }

func main() {
    var v1 string = "hello"
    var v2 uint = 2
    v3 := make(map[string]int)
    v4 := []string{("hello" + v1), v1, v1, ("foo bar" + "a")}
    v1 = "hello"
    if (false || false) {
        if false {
            for fi5 := 0; fi5 < 4; fi5++ {
                for fi6 := 0; fi6 < 13; fi6++ {
                    v1 = "hello"
                    _ = fi6
                }
                _ = fi5
            }
        } else {
            fmt.Println(v2, v2)
            var v7 int = 32768
            _ = v7
        }
        if true {
            v2 = (^(v2))
            v8 := v3[v1]
            v2 -= v2
            _ = v8
        } else {
            switch -627 {
            case 32768:
                for fi9 := 0; fi9 < 11; fi9++ {
                    if (true || true) {
                        fmt.Println(fi9, v2, v1)
                        for fi10 := 0; fi10 < 13; fi10++ {
                            if (!true) {
                                var v11 uint16 = 2
                                _ = v11
                            } else {
                                v1 = (v1 + (v1 + v1))
                            }
                            _ = fi10
                        }
                        v1 = ("a" + v1)
                    } else {
                        v1 = v1
                        if false {
                            if false {
                                v12 := GenS1{F0: false}
                                v2 -= v2
                                _ = v12
                            }
                            v1 = v1
                        } else {
                            _ = len(v4)
                        }
                    }
                    _ = fi9
                }
            case -103:
                v1 = v1
            case 0:
                var v13 int = 4294967295
                _ = v13
            default:
                var v14 uint32 = 2147483647
                _ = v14
            }
            v15 := v3["hello"]
            switch v1 {
            case "foo bar":
                _ = len(v4)
            }
            _ = v15
        }
        switch 651 {
        case -1:
            v1 = "foo bar"
        case 4294967295:
            v16 := GenS1{F0: (!false)}
            _ = v16
        default:
            for fi17 := 0; fi17 < 3; fi17++ {
                for fi18 := 0; fi18 < 8; fi18++ {
                    if true {
                        var v19 int8 = 127
                        var v20 int8 = (v19 + -17)
                        _ = v19
                        _ = v20
                    }
                    v4 = append(v4, v1)
                    _ = fi18
                }
                v21 := GenS2{F0: (fi17 % ((fi17) | 1)), F1: 1}
                v2 -= 264
                _ = fi17
                _ = v21
            }
        }
    } else {
        if (42 != 1) {
            if true {
                var v22 string = (v1 + v1)
                _ = v22
            } else {
                var v23 int16 = 0
                fmt.Println(v23)
                v24 := GenS2{F0: 32768, F1: 180}
                _ = v23
                _ = v24
            }
        }
        if false {
            if false {
                _ = len(v4)
                var v25 uint64 = 0
                _ = v25
            } else {
                v2 += 543
                if (!false) {
                    v1 = ("a" + v1)
                }
            }
            _ = len(v3)
        }
        v2 += 65535
    }
    for fi26 := 0; fi26 < 3; fi26++ {
        fmt.Println(fi26, v1)
        _ = fi26
    }
    if (3.14159 <= (1.5 / (-2.25 + -2.25))) {
        var v27 float32 = 1e10
        switch v1 {
        case "":
            if (0 >= 4294967295) {
                var v28 float64 = -2.25
                fmt.Println(v2, v27)
                _ = len(v3)
                _ = v28
            } else {
                if false {
                    v29 := GenS3{F0: v2, F1: v2, F2: 30}
                    var v30 bool = true
                    fmt.Println(v27)
                    _ = v29
                    _ = v30
                }
                switch 9223372036854775807 {
                case 0:
                    v1 = "hello"
                default:
                    var v31 uint64 = 18446744073709551615
                    _ = v31
                }
                v2 -= 948
            }
        case "\n\t":
            switch -533 {
            case -9223372036854775808:
                switch 564 {
                case -2147483648:
                    v3[(v1 + (v1 + v1))] = -673
                case 127:
                    fmt.Println(v2, v27)
                case 0:
                    for fi32 := 0; fi32 < 10; fi32++ {
                        if false {
                            _ = len(v3)
                        }
                        if len(v4) > 0 {
                            v4[0] = v1
                        }
                        _ = fi32
                    }
                default:
                    v2 = 2
                }
            default:
                if false {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
            }
        case "hello":
            v33 := GenS4{F0: 544, F1: -1, F2: "\n\t"}
            _ = v33
        }
        v2 = (v2 / (((v2 - (v2 - (v2 + (v2 >> ((v2) & 63)))))) | 1))
        _ = v27
    } else {
        v1 = v1
        var v34 uint64 = 4294967295
        _ = v34
    }
    var v35 int64 = -9223372036854775808
    v1 = (v1 + ((((((v1 + "hello") + "a") + ("\n\t" + ((v1 + v1) + v1))) + v1) + v1) + "\n\t"))
    _ = len(v4)
    if (42 == 2) {
        v36 := GenS1{F0: (583 >= 403)}
        var v37 string = v1
        _ = v36
        _ = v37
    } else {
        v1 = "a"
    }
    v3[("\n\t" + (v1 + "a"))] = 1
    v35 = 213
    v35 = v35
    if true {
        for fi38 := 0; fi38 < 11; fi38++ {
            fmt.Println(fi38, fi38)
            _ = fi38
        }
        var v39 uint8 = 127
        _ = v39
    } else {
        var v40 uint32 = uint32(v2)
        _ = v40
    }
    v35 = (v35 * 9223372036854775807)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
}
