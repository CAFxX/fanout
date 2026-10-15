package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int64 }
func genFn0(p0 string, p1 uint32, p2 uint8) int16 {
    fmt.Println(p0)
    switch 32768 {
    case -386:
        fmt.Println(p0, p1)
    case 128:
        v7 := GenS1{F0: 473, F1: 65535, F2: int64(p2)}
        _ = v7
    case 2147483647:
        v8 := GenS1{F0: 1, F1: 65535, F2: 42}
        _ = v8
    default:
        fmt.Println(p1)
    }
    if false {
        if false {
            p1 = (^(p1))
            for fi9 := 0; fi9 < 15; fi9++ {
                for fi10 := 0; fi10 < 4; fi10++ {
                    if false {
                        var v11 bool = true
                        var v12 int = fi10
                        _ = v11
                        _ = v12
                    } else {
                        var v13 bool = true
                        _ = v13
                    }
                    _ = fi10
                }
                var v14 uint64 = 895
                var v15 int16 = 0
                _ = fi9
                _ = v14
                _ = v15
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return -32768
}
func genFn1(p0 uint32, p1 uint) int {
    p1 = (p1 >> (((p1 | p1)) & 63))
    fmt.Println(p1, p0)
    var v16 uint = (p1 | p1)
    _ = p0
    _ = p1
    _ = v16
    return -78
}

func main() {
    var v1 uint16 = 611
    v2 := []string{("" + ("foo bar" + ("hello" + ("hello" + "hello")))), ("\n\t" + ("\n\t" + ("" + "a")))}
    var v3 uint64 = 4294967295
    var v4 float32 = 1e10
    var v5 string = "a"
    var v6 string = ((v5 + v5) + v5)
    if false {
        switch (((("" + v6) + v6) + "\n\t") + (v6 + v6)) {
        case "\n\t":
            _ = len(v2)
        default:
            for fi17 := 0; fi17 < 14; fi17++ {
                v5 = v6
                _ = fi17
            }
        }
        if (false && false) {
            v18 := GenS1{F0: 42, F1: v1, F2: 9223372036854775807}
            for fi19 := 0; fi19 < 17; fi19++ {
                var v20 int64 = 97
                var v21 uint16 = 65535
                v22 := GenS1{F0: fi19, F1: v21, F2: (v20 - v20)}
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
            }
            _ = v18
        }
        v6 = "\n\t"
    } else {
        var v23 float64 = -2.25
        _ = v23
    }
    if (true && (-881 == -1)) {
        var v24 int32 = -281
        var v25 int16 = 32767
        _ = v24
        _ = v25
    }
    switch 128 {
    case -2147483648:
        fmt.Println(v5)
    case 0:
        v4 += (float32(0.1) - 3.14159)
    }
    if true {
        var v26 uint = 4294967295
        for fi27 := 0; fi27 < 15; fi27++ {
            v28 := GenS1{F0: fi27, F1: (v1 >> (((v26 & 916)) & 63)), F2: 9223372036854775807}
            var v29 int = fi27
            if false {
                var v30 uint8 = 127
                for fi31 := 0; fi31 < 12; fi31++ {
                    v32 := GenS1{F0: -181, F1: v1, F2: 629}
                    v6 = "foo bar"
                    switch (((v6 + v5) + "a") + (v5 + v5)) {
                    case "foo bar":
                        if true {
                            fmt.Println(v6)
                            v33 := GenS1{F0: (fi31 - fi31), F1: (v1 << ((276) & 63)), F2: 42}
                            _ = v33
                        }
                    case "hello":
                        var v34 int8 = -30
                        _ = v34
                    }
                    _ = fi31
                    _ = v32
                }
                if len(v2) > 0 {
                    v2[0] = "foo bar"
                }
                _ = v30
            }
            _ = fi27
            _ = v28
            _ = v29
        }
        v35 := GenS1{F0: 4294967295, F1: v1, F2: 1}
        _ = v26
        _ = v35
    }
    v2 = append(v2, (v5 + "\n\t"))
    switch v6 {
    case "hello":
        var v36 uint16 = v1
        _ = v36
    default:
        for fi37 := 0; fi37 < 19; fi37++ {
            var v38 uint16 = (v1 & v1)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v39 int = -492
            _ = fi37
            _ = v38
            _ = v39
        }
    }
    v6 = ""
    for fi40 := 0; fi40 < 1; fi40++ {
        switch "\n\t" {
        case "hello":
            var v41 int32 = -98
            _ = v41
        default:
            v2 = append(v2, ("foo bar" + ((v6 + "\n\t") + (((v5 + "a") + ("a" + (v5 + "\n\t"))) + v5))))
        }
        var v42 int64 = 154
        v42 = (v42 * 9223372036854775807)
        _ = fi40
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
