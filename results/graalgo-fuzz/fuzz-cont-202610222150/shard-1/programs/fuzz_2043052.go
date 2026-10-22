package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 uint8 = 42
    v2 := []int{808, 32767, 32767}
    v3 := []int{4294967295, -629}
    var v4 int32 = -1
    var v5 uint32 = 4294967295
    var v6 uint16 = 32767
    var v7 bool = (false && false)
    var v8 bool = v7
    var v9 float32 = -2.25
    if ((v7 || v7) || (v1 < 171)) {
        var v10 int16 = -32768
        _ = v10
    } else {
        for fi11 := 0; fi11 < 20; fi11++ {
            if false {
                switch fi11 {
                case 136:
                    if v8 {
                        var v12 string = "a"
                        _ = v12
                    } else {
                        fmt.Println(v9)
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                case -808:
                    if false {
                        var v13 int16 = 1
                        v4 += v4
                        switch "" {
                        case "hello":
                            v7 = v8
                        default:
                            if len(v2) > 0 {
                                v2[0] = (-(fi11))
                            }
                        }
                        _ = v13
                    } else {
                        if true {
                            if v8 {
                                var v14 uint64 = 42
                                _ = v14
                            } else {
                                if false {
                                    for fi15 := 0; fi15 < 18; fi15++ {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        var v16 uint16 = (v6 + 282)
                                        _ = fi15
                                        _ = v16
                                    }
                                }
                                v1 += v1
                                for fi17 := 0; fi17 < 16; fi17++ {
                                    v18 := GenS1{F0: (fi17 | fi17)}
                                    _ = fi17
                                    _ = v18
                                }
                            }
                            var v19 int8 = 0
                            _ = v19
                        } else {
                            v9 += v9
                            if (665 != (v4 - (-(v4)))) {
                                if v7 {
                                    fmt.Println(v6, v5)
                                }
                                var v20 string = "a"
                                _ = v20
                            }
                            var v21 int16 = 32767
                            _ = v21
                        }
                        fmt.Println(fi11, v9)
                    }
                }
            }
            v7 = true
            _ = fi11
        }
        v6 += v6
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        for fi23 := 0; fi23 < 7; fi23++ {
            fmt.Println(v7, v6, v6)
            _ = fi23
        }
        _ = fi22
    }
    var v24 float64 = 1.5
    switch "" {
    case "a":
        var v25 uint = 548
        _ = v25
    case "\n\t":
        v4 += v4
    }
    if (v1 < (v1 | v1)) {
        for fi26 := 0; fi26 < 15; fi26++ {
            fmt.Println(v1)
            _ = len(v2)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi26
        }
        for fi27 := 0; fi27 < 3; fi27++ {
            var v28 int32 = 1
            var v29 int = (^(fi27))
            var v30 uint8 = v1
            _ = fi27
            _ = v28
            _ = v29
            _ = v30
        }
    }
    v31 := GenS2{F0: (v6 / (((v6 << ((0) & 63))) | 1))}
    var v32 int64 = -672
    for fi33 := 0; fi33 < 11; fi33++ {
        var v34 int = fi33
        _ = fi33
        _ = v34
    }
    var v35 bool = ("" > ("a" + "\n\t"))
    if (v1 > v1) {
        if len(v3) > 0 {
            v3[0] = 4294967295
        }
    }
    for fi36 := 0; fi36 < 7; fi36++ {
        switch 42 {
        case 127:
            var v37 float64 = v24
            _ = v37
        case 128:
            switch "a" {
            case "foo bar":
                var v38 string = "\n\t"
                _ = v38
            case "hello":
                v4 = (v4 ^ 42)
            default:
                if (("\n\t" + "") == "hello") {
                    v8 = (!v35)
                    if len(v2) > 0 {
                        v2[0] = 79
                    }
                    if v7 {
                        fmt.Println(v6, v5)
                        v6 += 500
                    }
                }
            }
        case -504:
            if true {
                if len(v3) > 0 {
                    v3[0] = -1
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v35 = v35
            }
        default:
            v1 += 141
        }
        v5 = (v5 & (v5 << ((2) & 63)))
        _ = fi36
    }
    switch -1 {
    case 2147483647:
        var v39 string = ("foo bar" + "")
        _ = v39
    default:
        v5 += (v5 >> ((637) & 63))
    }
    var v40 string = ((((("" + "\n\t") + ("" + "a")) + "hello") + "\n\t") + "foo bar")
    if ((684 >= -452) || true) {
        fmt.Println(v9)
    }
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v31
    _ = v32
    _ = v35
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
