package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 int16, p1 int8) uint8 {
    for fi7 := 0; fi7 < 4; fi7++ {
        for fi8 := 0; fi8 < 5; fi8++ {
            v9 := GenS1{F0: 1}
            var v10 uint = 65535
            v11 := []int{int(v10), (^(fi8))}
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
        for fi12 := 0; fi12 < 14; fi12++ {
            v13 := []int{506, fi7, (fi12 % ((fi7) | 1))}
            if len(v13) > 0 {
                v13[0] = fi12
            }
            v14 := GenS1{F0: 127}
            _ = fi12
            _ = v13
            _ = v14
        }
        v15 := GenS1{F0: 114}
        _ = fi7
        _ = v15
    }
    v16 := []int{535}
    p1 += (p1 - p1)
    _ = p0
    _ = p1
    _ = v16
    return 255
}
func genFn1(p0 uint8) uint64 {
    v17 := GenS1{F0: genFn0(-803, 0)}
    v18 := GenS1{F0: genFn0(-219, 42)}
    p0 = p0
    fmt.Println(p0)
    if ((((float64(float32(-2.25)) / float64(1.5)) < (-2.25 - (1e10 * 0.1))) || false) || ((1000 <= 1) || true)) {
        p0 = p0
        switch "foo bar" {
        case "hello":
            if (!false) {
                fmt.Println(p0)
                for fi19 := 0; fi19 < 18; fi19++ {
                    v20 := GenS1{F0: p0}
                    if true {
                        v21 := []int{fi19, (^(fi19))}
                        if len(v21) >= 2 {
                            v21 = v21[1:len(v21)]
                        }
                        _ = v21
                    }
                    _ = fi19
                    _ = v20
                }
                var v22 bool = false
                _ = v22
            }
        case "foo bar":
            fmt.Println(p0)
        case "\n\t":
            fmt.Println(p0)
        }
    } else {
        var v23 int = -9223372036854775808
        v24 := GenS1{F0: 42}
        var v25 int8 = 127
        _ = v23
        _ = v24
        _ = v25
    }
    _ = p0
    _ = v17
    _ = v18
    return 42
}

func main() {
    v1 := []uint32{65535, 194, 2147483647}
    var v2 int32 = 287
    v3 := make(map[int]bool)
    v4 := make(map[uint32]bool)
    var v5 int = -220
    var v6 uint32 = 791
    for fi26 := 0; fi26 < 14; fi26++ {
        v27 := GenS1{F0: 235}
        var v28 float32 = 0.0
        _ = fi26
        _ = v27
        _ = v28
    }
    v6 += (v6 % ((v6) | 1))
    v29 := GenS1{F0: 42}
    v5 = v5
    v2 = (-(v2))
    switch (("\n\t" + "foo bar") + "foo bar") {
    case "\n\t":
        var v30 float32 = 0.1
        _ = v30
    case "a":
        v31 := GenS1{F0: 92}
        _ = v31
    case "":
        v2 = (v2 << ((719) & 63))
    }
    var v32 uint64 = 816
    var v33 int8 = -80
    if (v5 != v5) {
        delete(v4, (v6 >> ((65535) & 63)))
    } else {
        for fi34 := 0; fi34 < 5; fi34++ {
            if true {
                if len(v1) > 0 {
                    v1[0] = v6
                }
                var v35 float32 = 3.14159
                _ = v35
            } else {
                v32 = genFn1(1)
            }
            _ = fi34
        }
        var v36 float64 = -2.25
        _ = v36
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
