package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 int64 = -1
    var v2 string = ""
    var v3 bool = (0 > 4294967295)
    v4 := make(map[uint32]string)
    var v5 bool = v3
    var v6 uint16 = 42
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 int16 = -1
        _ = fi7
        _ = v8
    }
    v1 = v1
    var v9 uint64 = 18446744073709551615
    fmt.Println(v5)
    if v5 {
        for fi10 := 0; fi10 < 8; fi10++ {
            v2 = "\n\t"
            _ = fi10
        }
    } else {
        v2 = v2
    }
    var v11 uint = 2
    switch v2 {
    case "\n\t":
        var v12 uint32 = 353
        _ = v12
    case "a":
        if ((3.14159 < 3.14159) && (!(-128 > 100))) {
            switch ((v2 + (v2 + ("a" + "hello"))) + "\n\t") {
            case "\n\t":
                v5 = (v9 != (v9 * v9))
            case "foo bar":
                if false {
                    var v13 string = ((((v2 + (v2 + "\n\t")) + v2) + "\n\t") + v2)
                    for fi14 := 0; fi14 < 17; fi14++ {
                        var v15 uint = v11
                        fmt.Println(v1, v2)
                        if v5 {
                            for fi16 := 0; fi16 < 2; fi16++ {
                                v17 := []int{(fi14 * 559)}
                                v1 += (v1 % (((v1 >> ((v15) & 63))) | 1))
                                if v3 {
                                    v13 = (("foo bar" + (("a" + "a") + (v13 + "hello"))) + v13)
                                    v4[2147483647] = "foo bar"
                                } else {
                                    fmt.Println(v15)
                                }
                                _ = fi16
                                _ = v17
                            }
                        } else {
                            v18 := GenS1{F0: 2147483648}
                            v3 = (!(!(719 > 0)))
                            _ = v18
                        }
                        _ = fi14
                        _ = v15
                    }
                    _ = v13
                }
            default:
                var v19 uint32 = 65535
                _ = v19
            }
        } else {
            fmt.Println(v9, v1, v5)
            v20 := []int{int(v6), 4294967295, -2147483648}
            _ = v20
        }
    case "foo bar":
        if v3 {
            var v21 int32 = 42
            _ = v21
        } else {
            fmt.Println(v2)
            v4[541] = v2
        }
    default:
        var v22 float32 = 1e10
        _ = v22
    }
    var v23 bool = false
    switch v2 {
    case "foo bar":
        if true {
            for fi24 := 0; fi24 < 9; fi24++ {
                v25 := []int{}
                delete(v4, 2147483648)
                _ = fi24
                _ = v25
            }
            var v26 uint16 = (v6 / (((v6 + v6)) | 1))
            _ = v26
        }
    default:
        var v27 uint = (v11 << (((v11 ^ v11)) & 63))
        _ = v27
    }
    v5 = true
    var v28 float32 = 1e10
    fmt.Println(v6, v11)
    var v29 bool = true
    v30 := []int{32767, 241}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
