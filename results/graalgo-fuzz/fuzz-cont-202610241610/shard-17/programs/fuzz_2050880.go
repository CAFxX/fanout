package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 uint8; F2 int16 }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[uint32]string)
    var v3 float64 = ((3.14159 / (((1e10 * 1.5) + 3.14159) * (-2.25 / ((((((1e10 * ((0.0 + 3.14159) + 1.5)) - 0.1) / 3.14159) - (1.5 * (float64(-2.25) - ((float64(-2.25) / 3.14159) / (1e10 - (3.14159 * 1.5)))))) / 0.1) + (3.14159 + 1.5))))) - 1.5)
    var v4 uint32 = 858
    var v5 float64 = float64(0.0)
    var v6 uint = 255
    v2[(v4 ^ v4)] = ((("\n\t" + "\n\t") + "a") + "hello")
    var v7 int64 = -444
    var v8 int8 = -24
    for fi9 := 0; fi9 < 10; fi9++ {
        if true {
            var v10 int = -9223372036854775808
            if false {
                delete(v2, 2147483648)
                for fi11 := 0; fi11 < 7; fi11++ {
                    v10 += (fi9 & 619)
                    _ = fi11
                }
            }
            v6 -= 65535
            _ = v10
        }
        var v12 string = "hello"
        v7 = (v7 * 1)
        _ = fi9
        _ = v12
    }
    var v13 uint32 = (^(v4))
    for fi14 := 0; fi14 < 19; fi14++ {
        v15 := GenS1{F0: (v7 - v7)}
        v4 = v13
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        v17 := []int{fi16, 1, fi16}
        if ((v1 >> ((v6) & 63)) == v1) {
            for fi18 := 0; fi18 < 17; fi18++ {
                v3 += (v5 * -2.25)
                _ = fi18
            }
            v4 = v4
        }
        for fi19 := 0; fi19 < 18; fi19++ {
            var v20 uint32 = 2147483648
            switch int(v8) {
            case -2147483648:
                var v21 bool = true
                _ = v21
            default:
                if true {
                    var v22 uint8 = 1
                    _ = v22
                } else {
                    v23 := GenS2{F0: (v13 / ((uint32(v6)) | 1)), F1: 128, F2: -144}
                    if ((fi16 - (fi16 >> (((v6 ^ 65535)) & 63))) != fi19) {
                        v7 = 866
                        v4 = v4
                        switch ("hello" + ("" + "\n\t")) {
                        case "hello":
                            var v24 int = (fi19 % ((1) | 1))
                            _ = v24
                        }
                        if true {
                            var v25 uint64 = (^(v1))
                            _ = v25
                        }
                        if true {
                            var v26 int8 = (-(v8))
                            var v27 int64 = (v7 << ((uint(v1)) & 63))
                            delete(v2, v20)
                            _ = v26
                            _ = v27
                        }
                    } else {
                        v28 := GenS2{F0: (v13 + (v4 >> ((v6) & 63))), F1: 108, F2: 172}
                        _ = v28
                    }
                    v29 := GenS3{F0: (fi16 | (fi19 / (((^(fi16))) | 1))), F1: 2}
                    if (((((true || (false || true)) && true) && ((1 <= v1) || true)) && false) && (1 > -1)) {
                        fmt.Println(v4)
                        v1 -= 18446744073709551615
                    }
                    if true {
                        v30 := GenS3{F0: fi19, F1: 255}
                        _ = v30
                    }
                    _ = v23
                    _ = v29
                }
            }
            v31 := GenS1{F0: v7}
            _ = fi19
            _ = v20
            _ = v31
        }
        _ = fi16
        _ = v17
    }
    switch "foo bar" {
    case "":
        if ((v6 * 18446744073709551615) == (v6 - (v6 + uint(v1)))) {
            fmt.Println(v1, v8)
            v32 := []int{}
            var v33 string = (((("a" + "hello") + "") + "\n\t") + (("hello" + "hello") + "hello"))
            _ = v32
            _ = v33
        } else {
            switch (("hello" + "") + (("a" + "foo bar") + (("\n\t" + "a") + "hello"))) {
            case "hello":
                v34 := GenS4{F0: (v1 >> (((v6 - v6)) & 63))}
                _ = v34
            default:
                v7 = (v7 << ((816) & 63))
            }
        }
    default:
        if (0 > -32768) {
            switch (("foo bar" + "") + "") {
            case "":
                fmt.Println(v6, v7, v1)
            }
        }
    }
    if (!(2147483647 <= -619)) {
        var v35 float64 = 3.14159
        _ = v35
    } else {
        fmt.Println(v5)
        var v36 int64 = (v7 / ((v7) | 1))
        _ = v36
    }
    for fi37 := 0; fi37 < 18; fi37++ {
        var v38 string = "\n\t"
        var v39 uint32 = (^(v4))
        _ = fi37
        _ = v38
        _ = v39
    }
    fmt.Println(v3, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
