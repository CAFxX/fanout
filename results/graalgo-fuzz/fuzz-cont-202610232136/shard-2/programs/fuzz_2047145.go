package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool }
func genFn0(p0 uint8, p1 int) uint64 {
    v4 := GenS1{F0: 32767}
    var v5 int32 = 201
    switch ("foo bar" + "a") {
    case "\n\t":
        for fi6 := 0; fi6 < 9; fi6++ {
            for fi7 := 0; fi7 < 11; fi7++ {
                v5 = (v5 << ((919) & 63))
                _ = fi7
            }
            _ = fi6
        }
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 18446744073709551615
}
func genFn1() uint16 {
    var v8 int16 = 32767
    if true {
        v9 := []int{int(v8)}
        _ = v9
    } else {
        v10 := GenS1{F0: 516}
        _ = v10
    }
    if (((v8 * v8) > v8) && (1 >= 0)) {
        if false {
            var v11 string = ("a" + "\n\t")
            _ = v11
        } else {
            v8 += v8
            v8 -= v8
        }
        fmt.Println(v8)
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        var v13 int32 = -2147483648
        var v14 float64 = (-2.25 + (float64(0.1) - (1e10 * 0.0)))
        fmt.Println(v8, v8, v13)
        _ = fi12
        _ = v13
        _ = v14
    }
    _ = v8
    return 328
}

func main() {
    var v1 float64 = (1e10 * ((0.1 * -2.25) + -2.25))
    v2 := make(map[uint32]int)
    var v3 bool = (!(101 != 42))
    fmt.Println(v3, v3)
    var v15 uint8 = 249
    for fi16 := 0; fi16 < 3; fi16++ {
        if (v1 < float64(float32(((float64(0.1) / 1.0) * 0.0)))) {
            var v17 string = ("\n\t" + "")
            v18 := GenS2{F0: v3}
            _ = v17
            _ = v18
        } else {
            for fi19 := 0; fi19 < 6; fi19++ {
                var v20 float64 = -2.25
                var v21 uint = 664
                _ = fi19
                _ = v20
                _ = v21
            }
            var v22 float64 = (-2.25 - ((v1 / -2.25) * float64(0.0)))
            _ = v22
        }
        _ = fi16
    }
    v15 = 181
    var v23 float32 = ((1.5 * 1e10) * 0.0)
    var v24 int = 1
    v25 := v2[4294967295]
    if (!v3) {
        for fi26 := 0; fi26 < 4; fi26++ {
            fmt.Println(v15, v1, v25)
            var v27 int16 = 795
            v28 := v2[1]
            _ = fi26
            _ = v27
            _ = v28
        }
        v23 += v23
        var v29 bool = false
        _ = v29
    }
    v30 := []int{}
    var v31 string = (("\n\t" + ("a" + "")) + "hello")
    var v32 uint16 = 1
    v33 := GenS2{F0: (!((!true) || (4294967295 < 759)))}
    if (true || v3) {
        v34 := GenS1{F0: v24}
        v32 -= (v32 << ((4294967295) & 63))
        var v35 int16 = 42
        _ = v34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
}
