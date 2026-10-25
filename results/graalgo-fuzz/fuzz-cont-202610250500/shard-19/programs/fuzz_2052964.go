package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 uint }
func genFn0() int32 {
    var v5 uint16 = 65535
    if true {
        fmt.Println(v5)
    } else {
        var v6 int = 1
        if true {
            v7 := []int{v6, v6, 128}
            v8 := GenS1{F0: 162, F1: 127, F2: (("hello" + "") + ("" + "a"))}
            _ = v7
            _ = v8
        }
        v9 := GenS2{F0: 2147483647}
        _ = v6
        _ = v9
    }
    _ = v5
    return -18
}

func main() {
    var v1 int = 32767
    var v2 uint = 65535
    v3 := []int{}
    v4 := make(map[uint32]int)
    v4[42] = (-(v1))
    v2 = v2
    _ = len(v4)
    switch ("hello" + "a") {
    case "\n\t":
        v10 := v4[0]
        _ = v10
    case "a":
        if true {
            if (847 <= 65535) {
                if false {
                    v11 := GenS2{F0: -2147483648}
                    var v12 uint = 863
                    v13 := GenS2{F0: 989}
                    _ = v11
                    _ = v12
                    _ = v13
                } else {
                    _ = len(v4)
                    switch v1 {
                    case 32767:
                        fmt.Println(v1, v2)
                    case -87:
                        delete(v4, uint32(v2))
                    }
                    var v14 int32 = 186
                    _ = v14
                }
            } else {
                if (!true) {
                    v2 = v2
                } else {
                    fmt.Println(v1, v1)
                    fmt.Println(v2)
                    v1 += (-(v1))
                }
                v15 := GenS2{F0: -928}
                var v16 int16 = 32767
                _ = v15
                _ = v16
            }
            v2 = v2
            if false {
                var v17 uint64 = 4294967295
                if ((42 == -623) && true) {
                    switch 61 {
                    case 42:
                        for fi18 := 0; fi18 < 2; fi18++ {
                            for fi19 := 0; fi19 < 9; fi19++ {
                                var v20 string = (("" + "a") + "\n\t")
                                v17 += (v17 / (((v17 ^ 42)) | 1))
                                _ = fi19
                                _ = v20
                            }
                            v1 += 2147483647
                            switch "a" {
                            case "foo bar":
                                v1 = v1
                            case "\n\t":
                                _ = len(v3)
                            default:
                                v21 := GenS3{F0: -170, F1: v2}
                                _ = v21
                            }
                            _ = fi18
                        }
                    case 128:
                        v22 := GenS3{F0: 229, F1: 18446744073709551615}
                        _ = v22
                    }
                }
                delete(v4, 2)
                _ = v17
            } else {
                for fi23 := 0; fi23 < 19; fi23++ {
                    v24 := GenS3{F0: (fi23 | v1), F1: (^(v2))}
                    _ = fi23
                    _ = v24
                }
                for fi25 := 0; fi25 < 8; fi25++ {
                    var v26 int32 = 159
                    var v27 string = "a"
                    _ = fi25
                    _ = v26
                    _ = v27
                }
                v1 -= (^(v1))
            }
        }
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v2 -= 434
    v2 -= v2
    var v28 string = "\n\t"
    var v29 bool = true
    switch (-(v1)) {
    case -2147483648:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case -9223372036854775808:
        fmt.Println(v1, v1, v1)
    default:
        var v30 uint8 = 42
        _ = v30
    }
    var v31 bool = (!v29)
    for fi32 := 0; fi32 < 20; fi32++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v33 int8 = 127
        _ = fi32
        _ = v33
    }
    v34 := GenS2{F0: 1}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v31 = (0 < -337)
    v35 := GenS2{F0: 2147483647}
    if (0 < 1000) {
        fmt.Println(v2, v28)
        fmt.Println(v2, v29)
        switch v28 {
        case "a":
            v1 -= -2147483648
        case "":
            v1 = (v1 | -857)
        case "foo bar":
            v31 = (!v31)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v34
    _ = v35
    _ = v4
}
