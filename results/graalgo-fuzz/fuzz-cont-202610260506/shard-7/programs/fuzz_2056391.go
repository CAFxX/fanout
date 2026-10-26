package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64 }
func genFn0() int8 {
    var v6 int32 = -643
    v7 := GenS1{F0: 658}
    _ = v6
    _ = v7
    return -121
}

func main() {
    var v1 uint64 = 4294967295
    var v2 string = "foo bar"
    var v3 bool = false
    v4 := []string{"a", v2, (v2 + (v2 + v2))}
    var v5 uint8 = 181
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 int = (fi8 % ((fi8) | 1))
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        if false {
            for fi11 := 0; fi11 < 13; fi11++ {
                var v12 uint64 = (v1 | v1)
                _ = fi11
                _ = v12
            }
            switch fi10 {
            case 2147483647:
                var v13 int64 = 42
                _ = v13
            }
        }
        _ = fi10
    }
    var v14 string = "foo bar"
    v4 = append(v4, v14)
    for fi15 := 0; fi15 < 10; fi15++ {
        var v16 int = (fi15 + fi15)
        var v17 string = v2
        _ = fi15
        _ = v16
        _ = v17
    }
    if v3 {
        var v18 bool = (0.1 != 3.14159)
        v19 := GenS1{F0: uint32(v1)}
        var v20 uint32 = 42
        _ = v18
        _ = v19
        _ = v20
    }
    if (!((41 >= 100) && (v3 && ((v1 + v1) >= 18446744073709551615)))) {
        var v21 int32 = -1
        _ = v21
    } else {
        var v22 uint64 = 0
        if (0 != 127) {
            if true {
                var v23 int32 = 456
                _ = v23
            } else {
                switch 1 {
                case 42:
                    var v24 int32 = 42
                    _ = v24
                default:
                    if v3 {
                        fmt.Println(v22, v14, v2)
                        if true {
                            switch -623 {
                            case 956:
                                v1 += v22
                            case -678:
                                v25 := GenS1{F0: 117}
                                _ = v25
                            }
                        }
                        var v26 uint16 = 2
                        _ = v26
                    } else {
                        var v27 uint32 = 4294967295
                        v4 = append(v4, ("hello" + ("hello" + "foo bar")))
                        fmt.Println(v14, v3, v14)
                        _ = v27
                    }
                }
                v2 = (v14 + v14)
                v28 := GenS2{F0: 100, F1: (v14 + ((((v2 + v2) + v14) + "\n\t") + v14))}
                _ = v28
            }
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v22
    }
    var v29 bool = v3
    v30 := GenS3{F0: 0}
    v3 = (0 > 2)
    var v31 uint16 = 42
    v2 = (v14 + ("" + v14))
    fmt.Println(v14)
    v1 += (^(v1))
    if v29 {
        var v32 uint64 = 4294967295
        if v29 {
            v32 = v32
        }
        _ = v32
    } else {
        v4 = append(v4, "hello")
    }
    if v29 {
        if (true && (!(!(1 >= -1)))) {
            fmt.Println(v29, v1)
        } else {
            for fi33 := 0; fi33 < 13; fi33++ {
                for fi34 := 0; fi34 < 13; fi34++ {
                    var v35 uint32 = 1
                    fmt.Println(v2, fi34)
                    fmt.Println(v29)
                    if false {
                        if len(v4) > 0 {
                            v4[0] = ("foo bar" + v2)
                        }
                        v4 = append(v4, ("foo bar" + "\n\t"))
                    }
                    _ = fi34
                    _ = v35
                }
                _ = fi33
            }
        }
    }
    switch v2 {
    case "hello":
        v4 = append(v4, v14)
    case "":
        switch v14 {
        case "a":
            if len(v4) > 0 {
                v4[0] = "foo bar"
            }
        case "hello":
            v4 = append(v4, (v14 + v14))
        case "\n\t":
            switch -257 {
            case -797:
                v36 := GenS2{F0: -111, F1: "a"}
                _ = v36
            case -1:
                if len(v4) > 0 {
                    v4[0] = ""
                }
            case 32767:
                for fi37 := 0; fi37 < 12; fi37++ {
                    if true {
                        var v38 uint64 = v1
                        for fi39 := 0; fi39 < 17; fi39++ {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            v40 := GenS2{F0: 27, F1: (v14 + ("hello" + (((("foo bar" + ((v14 + "a") + v2)) + v2) + v2) + v2)))}
                            fmt.Println(fi37)
                            _ = fi39
                            _ = v40
                        }
                        _ = v38
                    } else {
                        v41 := GenS1{F0: 65535}
                        fmt.Println(v14, v5)
                        _ = v41
                    }
                    var v42 int8 = genFn0()
                    _ = fi37
                    _ = v42
                }
            }
        default:
            var v43 float32 = -2.25
            _ = v43
        }
    }
    v44 := GenS4{F0: (v1 - 406)}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v44
    _ = v5
}
