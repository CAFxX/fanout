package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8; F1 uint64 }
func genFn0() int32 {
    var v4 int = 28
    for fi5 := 0; fi5 < 1; fi5++ {
        switch ("foo bar" + "a") {
        case "a":
            var v6 bool = ((fi5 << ((1) & 63)) >= v4)
            _ = v6
        case "hello":
            switch -2147483648 {
            case -193:
                var v7 float64 = ((3.14159 / 0.1) + float64((-2.25 / 1e10)))
                _ = v7
            default:
                v4 = v4
            }
        case "":
            var v8 uint32 = 183
            _ = v8
        }
        _ = fi5
    }
    v9 := GenS1{F0: 0, F1: 512}
    _ = v4
    _ = v9
    return 0
}
func genFn1(p0 uint8, p1 uint, p2 int16) uint {
    var v10 int = 128
    v11 := []int{v10, v10, (v10 % ((v10) | 1)), (v10 ^ v10)}
    p1 = 42
    p0 += (p0 >> ((p1) & 63))
    var v12 float32 = 1.5
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return (p1 & (p1 + 2))
}

func main() {
    var v1 string = "foo bar"
    var v2 uint = 18446744073709551615
    var v3 int8 = 108
    v2 = (v2 * genFn1(2, (v2 & 4294967295), 32767))
    v13 := []int{}
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    v14 := GenS1{F0: 0, F1: 814}
    fmt.Println(v2, v3)
    if true {
        var v15 float32 = (0.0 / 0.1)
        switch 2147483647 {
        case 42:
            fmt.Println(v15)
        case 127:
            if (false && false) {
                if len(v13) > 0 {
                    v13[0] = 42
                }
                fmt.Println(v3, v3)
            } else {
                for fi16 := 0; fi16 < 7; fi16++ {
                    if true {
                        var v17 uint64 = 2
                        var v18 uint = (v2 / ((v2) | 1))
                        _ = v17
                        _ = v18
                    } else {
                        var v19 uint64 = 42
                        _ = v19
                    }
                    switch (v1 + v1) {
                    case "hello":
                        _ = len(v13)
                    case "\n\t":
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                    case "":
                        switch fi16 {
                        case -9223372036854775808:
                            v3 -= v3
                        }
                    }
                    _ = fi16
                }
            }
        case 1:
            if ((1e10 - -2.25) == 0.0) {
                var v20 uint8 = 42
                v13 = append(v13, 2147483647)
                _ = v20
            }
        default:
            var v21 int8 = 127
            _ = v21
        }
        _ = v15
    } else {
        for fi22 := 0; fi22 < 6; fi22++ {
            v23 := GenS2{F0: genFn0()}
            v24 := GenS3{F0: 42}
            _ = fi22
            _ = v23
            _ = v24
        }
        v25 := GenS1{F0: 1, F1: 42}
        v26 := GenS4{F0: v3, F1: 18446744073709551615}
        _ = v25
        _ = v26
    }
    if (255 != 2) {
        switch v1 {
        case "a":
            v1 = v1
        case "\n\t":
            v27 := GenS2{F0: 1}
            _ = v27
        default:
            if (!(false && true)) {
                var v28 string = ((v1 + (((v1 + v1) + ("" + "\n\t")) + (v1 + "\n\t"))) + (("hello" + v1) + "foo bar"))
                fmt.Println(v28)
                _ = v28
            }
        }
        switch 127 {
        case 4294967295:
            fmt.Println(v2)
        }
        if (1.5 <= -2.25) {
            v1 = "a"
            if true {
                switch ("hello" + "hello") {
                case "foo bar":
                    var v29 uint = 2
                    _ = v29
                case "a":
                    v13 = append(v13, 0)
                case "\n\t":
                    v1 = v1
                }
            } else {
                if len(v13) > 0 {
                    v13[0] = 128
                }
                fmt.Println(v1)
            }
            switch v1 {
            case "hello":
                v13 = append(v13, 32767)
            case "\n\t":
                if len(v13) > 0 {
                    v13[0] = 32768
                }
            }
        } else {
            v1 = v1
        }
    }
    var v30 bool = true
    var v31 int32 = -1
    if len(v13) > 0 {
        v13[0] = 2147483647
    }
    var v32 float32 = (0.1 + float32(float64(float32((1e10 - (1e10 - (-2.25 * ((1e10 * 1.5) - (0.0 - 0.0)))))))))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
}
