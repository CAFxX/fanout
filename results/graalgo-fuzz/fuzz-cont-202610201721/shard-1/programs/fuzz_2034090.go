package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 uint = 42
    var v2 int = -1
    var v3 int64 = 1
    var v4 int64 = 1
    v5 := make(map[uint32]string)
    fmt.Println(v1, v3)
    var v6 uint8 = 0
    fmt.Println(v1)
    var v7 uint16 = 2
    var v8 uint8 = v6
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 int = (v2 & -9223372036854775808)
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: 42}
    for fi12 := 0; fi12 < 15; fi12++ {
        var v13 uint64 = 4294967295
        _ = fi12
        _ = v13
    }
    v8 -= v6
    v6 = (v8 % (((^(v6))) | 1))
    var v14 uint32 = 2147483648
    var v15 int64 = v3
    var v16 int8 = 127
    var v17 uint8 = v6
    if false {
        var v18 uint64 = 0
        var v19 string = "a"
        var v20 uint8 = v17
        _ = v18
        _ = v19
        _ = v20
    } else {
        v21 := GenS1{F0: (v14 | (^(v14)))}
        if false {
            if true {
                for fi22 := 0; fi22 < 18; fi22++ {
                    fmt.Println(fi22, v6, v17)
                    v1 = 929
                    fmt.Println(v2, v3)
                    _ = fi22
                }
                fmt.Println(v1, v2)
            }
            if true {
                fmt.Println(v7, v2, v3)
                var v23 uint8 = (^(v8))
                _ = v23
            }
        } else {
            switch "a" {
            case "\n\t":
                v24 := GenS1{F0: 0}
                _ = v24
            case "":
                var v25 string = "a"
                _ = v25
            case "a":
                var v26 float32 = ((0.0 / 0.1) * (0.1 - (-2.25 * ((1e10 / 3.14159) + -2.25))))
                _ = v26
            default:
                var v27 uint32 = (v14 - uint32(v1))
                _ = v27
            }
        }
        _ = v21
    }
    var v28 float64 = 3.14159
    if ((1000 <= 32767) || true) {
        v15 += (v4 ^ v3)
        if ((v6 << ((v1) & 63)) == uint8(v7)) {
            if true {
                switch (("hello" + (((("hello" + "foo bar") + "a") + "a") + ("\n\t" + (("" + "hello") + "hello")))) + (((("" + "\n\t") + "a") + "hello") + "foo bar")) {
                case "hello":
                    v29 := []int{v2, (-(v2))}
                    _ = v29
                case "":
                    if ((1 <= 1) || false) {
                        v7 -= (v7 - (v7 + 255))
                        v15 += -9223372036854775808
                    }
                case "a":
                    v30 := GenS1{F0: (^(v14))}
                    _ = v30
                }
                if ((1.5 / (0.1 + float32(((((-2.25 * 3.14159) / v28) + (3.14159 - v28)) - (((v28 / 3.14159) / -2.25) - 1e10))))) >= float32(float64((0.1 * 3.14159)))) {
                    for fi31 := 0; fi31 < 3; fi31++ {
                        for fi32 := 0; fi32 < 16; fi32++ {
                            v15 = (^(v4))
                            v1 -= (^(v1))
                            _ = fi32
                        }
                        v16 -= v16
                        var v33 float32 = 1.5
                        _ = fi31
                        _ = v33
                    }
                    v34 := []int{v2}
                    _ = v34
                }
                v35 := []int{}
                _ = v35
            }
            v36 := GenS1{F0: 2}
            _ = v36
        } else {
            v37 := v5[(v14 - (v14 >> (((v1 + v1)) & 63)))]
            _ = v37
        }
        var v38 bool = (false && (int64(v14) >= v15))
        _ = v38
    } else {
        v8 = 200
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
