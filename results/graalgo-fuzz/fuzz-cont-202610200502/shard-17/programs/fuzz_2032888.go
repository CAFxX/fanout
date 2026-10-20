package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint64 }
type GenS2 struct { F0 int64; F1 uint32 }
func genFn0(p0 uint32, p1 uint32, p2 uint32) uint8 {
    v7 := []int{2147483647}
    fmt.Println(p1)
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    var v8 uint16 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 int16 = 32767
    var v2 int32 = -871
    var v3 int16 = -32768
    var v4 int32 = 2147483647
    v5 := make(map[int]bool)
    v6 := []int{2147483647, 209, 127}
    var v9 float32 = -2.25
    v2 -= 654
    fmt.Println(v3, v4, v4)
    if ((("a" + "\n\t") + "a") >= "\n\t") {
        if (612 != 885) {
            for fi10 := 0; fi10 < 2; fi10++ {
                v11 := GenS1{F0: 42, F1: -1, F2: 638}
                for fi12 := 0; fi12 < 10; fi12++ {
                    var v13 uint = 255
                    v14 := GenS1{F0: v13, F1: 42, F2: 18446744073709551615}
                    fmt.Println(v1)
                    v3 += (v3 | (v3 >> ((42) & 63)))
                    var v15 int64 = 1
                    _ = fi12
                    _ = v13
                    _ = v14
                    _ = v15
                }
                if (false || true) {
                    v5[(^(fi10))] = (false || (18446744073709551615 == 324))
                }
                _ = fi10
                _ = v11
            }
            v16 := GenS2{F0: 947, F1: 2147483648}
            _ = v16
        } else {
            v3 -= (v1 + -32768)
        }
        if false {
            v4 = 412
        } else {
            var v17 bool = ((486 >= 2147483647) && (true || true))
            v18 := GenS2{F0: 39, F1: 42}
            _ = v17
            _ = v18
        }
        switch ((("foo bar" + (("\n\t" + "foo bar") + "a")) + "") + "a") {
        case "foo bar":
            v2 += 1
        case "hello":
            var v19 int16 = v3
            _ = v19
        case "\n\t":
            v1 -= (v3 >> ((348) & 63))
        }
    } else {
        var v20 float32 = 0.0
        if true {
            var v21 float32 = (((3.14159 * float32(-2.25)) * (1.5 - float32((float64(v9) / (-2.25 / 1e10))))) / (-2.25 * v20))
            _ = v21
        }
        fmt.Println(v2)
        _ = v20
    }
    switch -9223372036854775808 {
    case -9223372036854775808:
        var v22 int64 = -164
        _ = v22
    case 9223372036854775807:
        if (v3 > v1) {
            for fi23 := 0; fi23 < 11; fi23++ {
                var v24 string = ((("" + ("hello" + "")) + "") + ("a" + ((("" + "\n\t") + "") + "")))
                _ = fi23
                _ = v24
            }
            var v25 bool = (29 <= 42)
            if len(v6) > 0 {
                v6[0] = 4294967295
            }
            _ = v25
        }
    case -47:
        for fi26 := 0; fi26 < 7; fi26++ {
            v27 := v5[1]
            _ = fi26
            _ = v27
        }
    }
    v28 := v5[1]
    if v28 {
        var v29 int64 = -9223372036854775808
        var v30 int16 = v3
        _ = v29
        _ = v30
    }
    v31 := GenS2{F0: -924, F1: 0}
    switch ((("hello" + "\n\t") + "hello") + "a") {
    case "foo bar":
        delete(v5, 1)
    case "\n\t":
        v5[-1] = ((-(v3)) <= 1000)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
