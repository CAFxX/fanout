package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 int32; F2 uint8 }
func genFn0(p0 uint8, p1 uint8, p2 int16) uint64 {
    v6 := []int{-1}
    for fi7 := 0; fi7 < 19; fi7++ {
        p1 = (^(p0))
        var v8 bool = (1 == 0)
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 891
}
func genFn1(p0 uint16) uint32 {
    v9 := GenS1{F0: 450}
    var v10 int8 = -1
    p0 -= 2
    switch ("\n\t" + "") {
    case "a":
        p0 += 255
    case "foo bar":
        p0 = p0
    default:
        v11 := []int{-2147483648, 0, -2147483648, 9223372036854775807}
        _ = v11
    }
    _ = p0
    _ = v10
    _ = v9
    return 2
}

func main() {
    var v1 uint = 697
    v2 := make(map[int]string)
    var v3 int32 = 991
    var v4 uint32 = 472
    var v5 uint16 = 32767
    var v12 uint8 = 127
    v13 := GenS1{F0: v5}
    var v14 uint32 = v4
    fmt.Println(v1, v1, v12)
    for fi15 := 0; fi15 < 5; fi15++ {
        for fi16 := 0; fi16 < 3; fi16++ {
            delete(v2, (^(fi16)))
            var v17 int32 = (v3 + v3)
            _ = fi16
            _ = v17
        }
        for fi18 := 0; fi18 < 6; fi18++ {
            if ((v1 % ((v1) | 1)) < v1) {
                var v19 string = "a"
                _ = v19
            } else {
                var v20 bool = false
                v21 := GenS1{F0: (v5 * v5)}
                _ = v20
                _ = v21
            }
            _ = fi18
        }
        if (-1 >= (fi15 ^ (-(fi15)))) {
            if false {
                v22 := GenS1{F0: 32767}
                _ = v22
            } else {
                v14 = genFn1(177)
            }
            fmt.Println(v14, v3)
        } else {
            var v23 bool = (true && (v5 < v5))
            _ = v23
        }
        _ = fi15
    }
    if true {
        fmt.Println(v14, v14, v3)
        switch 0 {
        case 1:
            switch -9223372036854775808 {
            case 128:
                var v24 float32 = (0.0 + (((float32((0.1 - 0.1)) * 0.1) / 1.0) / float32(0.1)))
                _ = v24
            case -841:
                v25 := GenS2{F0: true, F1: 1, F2: (v12 >> ((v1) & 63))}
                _ = v25
            case 4294967295:
                var v26 int32 = (-(v3))
                _ = v26
            }
        case -9223372036854775808:
            v27 := GenS1{F0: v5}
            _ = v27
        case -2147483648:
            v4 -= (v14 >> (((v1 << (((v1 | v1)) & 63))) & 63))
        }
        var v28 float64 = (0.1 / (((((1e10 + 1.5) / float64((0.1 / 1.0))) + (((1.5 * 0.1) * (-2.25 - 0.0)) / ((3.14159 / (-2.25 - float64(0.1))) + (0.0 - float64(3.14159))))) / (1e10 - 1.5)) - float64(((float32((0.1 / 1e10)) * 0.1) + (0.1 / -2.25)))))
        _ = v28
    } else {
        fmt.Println(v4, v1, v14)
    }
    for fi29 := 0; fi29 < 3; fi29++ {
        v3 += v3
        v30 := GenS1{F0: 382}
        _ = fi29
        _ = v30
    }
    for fi31 := 0; fi31 < 12; fi31++ {
        v3 += -1
        switch ("hello" + "") {
        case "a":
            v32 := GenS1{F0: (v5 - v5)}
            _ = v32
        case "hello":
            v33 := []int{fi31, fi31, fi31}
            _ = v33
        }
        _ = fi31
    }
    v12 -= (v12 ^ (v12 | (^(v12))))
    v14 = (v4 - (v4 >> (((v1 ^ uint(v12))) & 63)))
    v5 += (v5 | (v5 >> (((v1 - (v1 ^ uint(v5)))) & 63)))
    var v34 uint32 = v14
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
