package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 uint16; F2 string }
type GenS4 struct { F0 int8; F1 int8; F2 uint8 }
type GenS5 struct { F0 uint; F1 int8; F2 int16 }
func genFn0(p0 uint64, p1 uint8, p2 int) uint8 {
    v5 := []int{p2}
    p0 += p0
    var v6 uint16 = 699
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return p1
}

func main() {
    v1 := make(map[string]string)
    v2 := []int{-2147483648}
    v3 := []int{}
    v4 := []int{-554, 2147483647, -2147483648}
    var v7 int16 = -32768
    v8 := GenS1{F0: 2, F1: false, F2: -636}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1["a"] = ("\n\t" + ("a" + "a"))
    fmt.Println(v7)
    var v9 int8 = 35
    v10 := GenS1{F0: 0, F1: (2147483648 >= 42), F2: -774}
    switch 9223372036854775807 {
    case 2147483647:
        for fi11 := 0; fi11 < 13; fi11++ {
            if (277 >= 2147483647) {
                switch "foo bar" {
                case "":
                    _ = len(v3)
                case "hello":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
            }
            v12 := GenS1{F0: 488, F1: false, F2: 3}
            _ = fi11
            _ = v12
        }
    case -1:
        fmt.Println(v7)
    case -874:
        v13 := GenS1{F0: 4294967295, F1: (!(true && true)), F2: -278}
        _ = v13
    default:
        switch "\n\t" {
        case "":
            if false {
                v14 := GenS2{F0: ("foo bar" + ("foo bar" + "hello"))}
                _ = v14
            }
        case "foo bar":
            v7 -= v7
        default:
            v15 := GenS3{F0: 109, F1: 65535, F2: ("hello" + ("\n\t" + ((("hello" + "hello") + (("\n\t" + ("foo bar" + "foo bar")) + "hello")) + "foo bar")))}
            _ = v15
        }
    }
    v9 = v9
    if true {
        v16 := GenS4{F0: 4, F1: (v9 & v9), F2: 42}
        _ = v16
    } else {
        switch 1 {
        case 42:
            for fi17 := 0; fi17 < 14; fi17++ {
                var v18 uint8 = 0
                _ = fi17
                _ = v18
            }
        case 347:
            var v19 uint64 = 1
            _ = v19
        default:
            v20 := GenS2{F0: ("" + "foo bar")}
            _ = v20
        }
        var v21 uint16 = 1
        _ = v21
    }
    if (!(!false)) {
        var v22 int = -959
        var v23 uint32 = 2
        _ = v22
        _ = v23
    }
    var v24 int8 = (v9 - (v9 + (v9 | (v9 / (((v9 & v9)) | 1)))))
    switch "\n\t" {
    case "foo bar":
        v9 = (v24 - 0)
    case "a":
        switch "foo bar" {
        case "a":
            v25 := GenS1{F0: 0, F1: true, F2: 4294967295}
            _ = v25
        case "hello":
            for fi26 := 0; fi26 < 20; fi26++ {
                v27 := GenS5{F0: 42, F1: (v24 * v24), F2: -332}
                switch (fi26 << ((4294967295) & 63)) {
                case -365:
                    delete(v1, "hello")
                case 9223372036854775807:
                    if ((fi26 > fi26) && (127 > 1)) {
                        var v28 int32 = -666
                        v7 -= v7
                        _ = v28
                    } else {
                        switch fi26 {
                        case 32768:
                            v4 = append(v4, fi26)
                        default:
                            if (-45 <= 100) {
                                var v29 string = "hello"
                                v30 := GenS2{F0: (((v29 + "") + "\n\t") + v29)}
                                _ = v29
                                _ = v30
                            }
                        }
                        v2 = append(v2, -137)
                        for fi31 := 0; fi31 < 5; fi31++ {
                            v7 += (v7 / ((v7) | 1))
                            v24 = v24
                            var v32 int64 = 42
                            _ = fi31
                            _ = v32
                        }
                    }
                case 32768:
                    v33 := GenS5{F0: 1000, F1: 42, F2: v7}
                    _ = v33
                default:
                    _ = len(v3)
                }
                _ = fi26
                _ = v27
            }
        case "foo bar":
            var v34 int16 = (^(v7))
            _ = v34
        }
    case "\n\t":
        v2 = append(v2, 4294967295)
    }
    v35 := GenS4{F0: (v9 ^ 127), F1: (v24 | (v9 & (v24 & 127))), F2: 237}
    fmt.Println(v24)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v3
    _ = v35
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
