package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 int }

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[uint32]int)
    v3 := []string{"hello", "\n\t", "hello", ("" + ("hello" + "a"))}
    if false {
        for fi4 := 0; fi4 < 19; fi4++ {
            v1 = v1
            _ = fi4
        }
        _ = len(v2)
    }
    if ((0.0 + (0.1 * -2.25)) > ((float32((-2.25 * ((3.14159 / (float64((-2.25 - -2.25)) - ((0.0 - float64((0.0 * 0.1))) + ((float64((1.5 / 0.1)) + -2.25) / 3.14159)))) * (((float64(float32(1e10)) * -2.25) - 0.0) + 3.14159)))) - 1.5) / 0.1)) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v1 = v1
    }
    if (0 > 1000) {
        v1 -= v1
        if (true || true) {
            v1 += v1
            var v5 int32 = 88
            _ = v5
        } else {
            v1 += (v1 << ((65535) & 63))
            if true {
                var v6 string = "\n\t"
                if (2 <= 83) {
                    _ = len(v3)
                    v1 += v1
                    for fi7 := 0; fi7 < 12; fi7++ {
                        _ = len(v2)
                        v6 = ("\n\t" + (v6 + "foo bar"))
                        _ = fi7
                    }
                }
                if ((v1 | -943) <= 1) {
                    var v8 int64 = -1
                    _ = v8
                }
                _ = v6
            } else {
                v1 = 0
                v1 = 0
                delete(v2, 2147483648)
            }
        }
        var v9 int64 = v1
        _ = v9
    } else {
        v10 := GenS1{F0: ""}
        if len(v3) > 0 {
            v3[0] = "\n\t"
        }
        _ = v10
    }
    delete(v2, 1)
    v2[42] = 2147483647
    for fi11 := 0; fi11 < 12; fi11++ {
        for fi12 := 0; fi12 < 4; fi12++ {
            if true {
                v13 := GenS1{F0: "a"}
                v3 = append(v3, "")
                v1 -= v1
                switch "hello" {
                case "a":
                    if ((fi11 - -1) == fi11) {
                        v14 := v2[4294967295]
                        var v15 float64 = 1e10
                        var v16 uint8 = 1
                        _ = v14
                        _ = v15
                        _ = v16
                    } else {
                        v17 := GenS1{F0: ("a" + "hello")}
                        var v18 bool = false
                        _ = v17
                        _ = v18
                    }
                }
                _ = v13
            }
            var v19 uint64 = 2
            v20 := GenS2{F0: 42, F1: 128}
            _ = fi12
            _ = v19
            _ = v20
        }
        v21 := GenS2{F0: 0, F1: fi11}
        v22 := GenS2{F0: 774, F1: fi11}
        _ = fi11
        _ = v21
        _ = v22
    }
    if false {
        v1 += v1
    }
    var v23 uint8 = 0
    v24 := GenS1{F0: "a"}
    var v25 uint64 = 18446744073709551615
    if false {
        switch 32767 {
        case -2147483648:
            var v26 float32 = ((1e10 + ((float32(-2.25) * float32(float64(1.5))) / 0.1)) / 0.1)
            _ = v26
        case -580:
            v25 += (v25 ^ (v25 - v25))
        case 9223372036854775807:
            for fi27 := 0; fi27 < 1; fi27++ {
                var v28 string = "\n\t"
                _ = fi27
                _ = v28
            }
        }
        fmt.Println(v23, v25)
    }
    _ = len(v3)
    v29 := GenS1{F0: "foo bar"}
    var v30 uint64 = v25
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v30
}
