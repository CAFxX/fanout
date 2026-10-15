package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 uint64 }
func genFn0() uint32 {
    v4 := GenS1{F0: 103}
    v5 := []int{2147483647, -1, 9223372036854775807}
    var v6 float32 = (1.5 / 1.0)
    _ = v4
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 int16) uint64 {
    fmt.Println(p0)
    v7 := []int{-2147483648}
    _ = p0
    _ = v7
    return 18446744073709551615
}

func main() {
    var v1 uint64 = 1
    var v2 int32 = -1
    var v3 uint32 = 819
    switch 32767 {
    case 1:
        v8 := []int{127}
        _ = v8
    }
    v1 += (v1 - v1)
    v9 := GenS2{F0: 42}
    if ((v1 << ((uint(v3)) & 63)) < v1) {
        for fi10 := 0; fi10 < 20; fi10++ {
            fmt.Println(v1, v1, v2)
            v11 := []int{}
            _ = fi10
            _ = v11
        }
    }
    switch ("foo bar" + ("hello" + "hello")) {
    case "\n\t":
        for fi12 := 0; fi12 < 16; fi12++ {
            v2 += v2
            if false {
                var v13 int32 = (v2 >> ((255) & 63))
                _ = v13
            }
            _ = fi12
        }
    case "a":
        if false {
            switch -1 {
            case -2147483648:
                var v14 uint64 = v1
                _ = v14
            }
            var v15 uint8 = 151
            _ = v15
        } else {
            v16 := GenS1{F0: uint8(v3)}
            v2 += (v2 << ((2) & 63))
            _ = v16
        }
    case "foo bar":
        switch "hello" {
        case "hello":
            var v17 uint32 = v3
            _ = v17
        }
    default:
        if true {
            v18 := GenS2{F0: 186}
            _ = v18
        }
    }
    v19 := []int{916, 338, 289}
    var v20 uint16 = 42
    if true {
        if len(v19) > 0 {
            v19[0] = 32768
        }
    }
    for fi21 := 0; fi21 < 18; fi21++ {
        switch fi21 {
        case 128:
            if len(v19) >= 2 {
                v19 = v19[1:len(v19)]
            }
        case 9223372036854775807:
            _ = len(v19)
        }
        if true {
            var v22 int32 = (v2 | -1)
            _ = len(v19)
            _ = v22
        }
        switch fi21 {
        case -1:
            if ((v1 >= genFn1(-32768)) && true) {
                for fi23 := 0; fi23 < 13; fi23++ {
                    fmt.Println(v1, v3)
                    fmt.Println(v2, fi23)
                    var v24 uint32 = 4294967295
                    _ = fi23
                    _ = v24
                }
            } else {
                if (1 > 9223372036854775807) {
                    v25 := GenS2{F0: 141}
                    _ = v25
                }
                v26 := GenS3{F0: (v3 << ((669) & 63)), F1: v1}
                _ = v26
            }
        default:
            for fi27 := 0; fi27 < 17; fi27++ {
                v19 = append(v19, (fi21 * fi27))
                _ = fi27
            }
        }
        _ = fi21
    }
    var v28 uint16 = (v20 / (((v20 * (^(v20)))) | 1))
    var v29 uint16 = (v28 / ((42) | 1))
    var v30 uint16 = (v28 * v28)
    switch 4294967295 {
    case -9223372036854775808:
        switch -2147483648 {
        case -472:
            switch 32767 {
            case 32768:
                var v31 uint8 = 2
                _ = v31
            case 654:
                fmt.Println(v20, v20, v1)
            case 203:
                var v32 uint32 = (v3 * 929)
                _ = v32
            }
        default:
            if (!(42 > 157)) {
                v33 := GenS2{F0: 138}
                v29 -= 119
                fmt.Println(v30)
                _ = v33
            } else {
                v19 = append(v19, 4294967295)
            }
        }
    default:
        var v34 uint = 0
        _ = v34
    }
    var v35 uint16 = v29
    switch -16 {
    case 115:
        v3 -= 65535
    case 42:
        if len(v19) >= 2 {
            v19 = v19[1:len(v19)]
        }
    case 0:
        if len(v19) >= 2 {
            v19 = v19[1:len(v19)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v35
    _ = v9
}
