package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint }
func genFn0() string {
    v4 := GenS1{F0: -1}
    var v5 uint = 18446744073709551615
    var v6 bool = ((1.5 / (0.0 + -2.25)) == (0.0 - 0.1))
    _ = v4
    _ = v5
    _ = v6
    return "hello"
}

func main() {
    var v1 float32 = -2.25
    v2 := make(map[uint32]int)
    var v3 int32 = -1
    var v7 uint = 4294967295
    v7 -= 119
    delete(v2, 0)
    v3 += v3
    v7 += v7
    if false {
        switch 128 {
        case 128:
            v1 -= (((v1 * ((-2.25 - v1) + v1)) + v1) * v1)
        case -2147483648:
            fmt.Println(v1, v1, v7)
        }
        var v8 uint = v7
        _ = v8
    } else {
        var v9 int8 = 127
        if false {
            for fi10 := 0; fi10 < 19; fi10++ {
                v11 := []int{(fi10 + (fi10 - 0)), fi10}
                _ = fi10
                _ = v11
            }
            v9 = int8(v7)
        }
        _ = v9
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        var v13 uint64 = 0
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 13; fi14++ {
        switch ("" + (("" + (((("\n\t" + ((("\n\t" + "hello") + "\n\t") + "\n\t")) + "hello") + "") + "a")) + ((("foo bar" + "") + ("" + ("foo bar" + "\n\t"))) + ("\n\t" + ("\n\t" + "\n\t"))))) {
        case "a":
            for fi15 := 0; fi15 < 8; fi15++ {
                switch fi14 {
                case -180:
                    v16 := GenS2{F0: 2147483647, F1: 2, F2: (v7 << (((v7 * 2)) & 63))}
                    _ = v16
                }
                _ = fi15
            }
        case "foo bar":
            switch fi14 {
            case -9223372036854775808:
                var v17 int16 = 251
                _ = v17
            case 4294967295:
                v3 = v3
            }
        case "":
            if (65535 != 32767) {
                v18 := []int{fi14}
                _ = v18
            } else {
                v19 := []int{}
                v3 -= -2147483648
                _ = v19
            }
        }
        var v20 uint = (v7 << ((v7) & 63))
        _ = fi14
        _ = v20
    }
    var v21 float64 = (0.1 + (3.14159 + -2.25))
    fmt.Println(v3, v3)
    _ = len(v2)
    var v22 uint64 = 765
    if (!false) {
        v1 -= v1
    } else {
        var v23 int = -2147483648
        v23 += 772
        _ = v23
    }
    if true {
        var v24 string = ("a" + "foo bar")
        _ = v24
    } else {
        v2[1] = 32767
        v1 = 0.1
    }
    for fi25 := 0; fi25 < 8; fi25++ {
        var v26 uint32 = 2147483648
        var v27 int8 = -128
        _ = fi25
        _ = v26
        _ = v27
    }
    var v28 uint = v7
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v28
    _ = v3
    _ = v7
}
