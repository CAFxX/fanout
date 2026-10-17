package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 uint32 = 0
    v2 := []string{(("a" + ("" + "foo bar")) + (((("a" + "hello") + "\n\t") + (("foo bar" + ("hello" + ("a" + "a"))) + "\n\t")) + "foo bar")), (("hello" + "\n\t") + (("foo bar" + "") + "a")), "", "\n\t"}
    v3 := make(map[string]int)
    var v4 float32 = (1e10 - 3.14159)
    if true {
        for fi5 := 0; fi5 < 18; fi5++ {
            v6 := GenS1{F0: 61}
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi5
            _ = v6
        }
        if (!false) {
            if true {
                var v7 float32 = 0.0
                v7 = v4
                if (1 != 1) {
                    for fi8 := 0; fi8 < 16; fi8++ {
                        var v9 int32 = -239
                        if false {
                            var v10 uint64 = 2
                            if false {
                                for fi11 := 0; fi11 < 17; fi11++ {
                                    fmt.Println(v7, v7, v7)
                                    _ = fi11
                                }
                                var v12 uint32 = (v1 * 42)
                                fmt.Println(fi8)
                                _ = v12
                            } else {
                                v9 = v9
                            }
                            _ = v10
                        }
                        _ = fi8
                        _ = v9
                    }
                    v1 += v1
                }
                _ = v7
            } else {
                delete(v3, "")
                v13 := GenS1{F0: 105}
                switch ("" + (("foo bar" + ("hello" + (("foo bar" + "\n\t") + "\n\t"))) + "foo bar")) {
                case "\n\t":
                    var v14 uint8 = 186
                    _ = v14
                default:
                    v4 = (v4 / 1.0)
                }
                _ = v13
            }
            switch -92 {
            case 622:
                switch "a" {
                case "hello":
                    v15 := GenS1{F0: 100}
                    _ = v15
                case "a":
                    v1 += 4294967295
                case "foo bar":
                    var v16 uint = 0
                    _ = v16
                }
            }
        } else {
            if true {
                var v17 bool = true
                for fi18 := 0; fi18 < 16; fi18++ {
                    if false {
                        fmt.Println(v17)
                        var v19 int64 = -610
                        fmt.Println(v19, v1, v1)
                        _ = v19
                    }
                    var v20 uint64 = 813
                    fmt.Println(v17, v17, fi18)
                    _ = fi18
                    _ = v20
                }
                _ = v17
            }
        }
        v21 := GenS1{F0: -1}
        _ = v21
    }
    var v22 int = 0
    switch v22 {
    case 9223372036854775807:
        v1 -= v1
    case 520:
        var v23 uint16 = 255
        _ = v23
    case 38:
        v24 := v3["hello"]
        _ = v24
    default:
        var v25 uint32 = v1
        _ = v25
    }
    v2 = append(v2, "\n\t")
    if (((1e10 - float32((3.14159 + 0.0))) != (1.5 / -2.25)) && (0 > -32768)) {
        v26 := GenS1{F0: 1}
        _ = len(v2)
        switch "" {
        case "foo bar":
            var v27 int32 = -904
            _ = v27
        }
        _ = v26
    } else {
        switch (v22 + (v22 / ((v22) | 1))) {
        case 2147483647:
            v4 = v4
        case -2147483648:
            v22 += v22
        case 208:
            var v28 int32 = -1
            _ = v28
        }
        v1 += (^(v1))
    }
    v29 := GenS1{F0: 1}
    var v30 uint8 = 255
    var v31 int8 = 100
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
