package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 int8; F2 uint32 }
type GenS4 struct { F0 uint8; F1 int16 }
func genFn0(p0 uint32, p1 uint64, p2 string) bool {
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 bool = ((p0 << ((0) & 63)) == 0)
        v9 := GenS1{F0: p1, F1: 100}
        p1 -= 18446744073709551615
        _ = fi7
        _ = v8
        _ = v9
    }
    fmt.Println(p0, p0)
    p1 = p1
    var v10 float32 = 0.1
    p0 = 623
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return true
}
func genFn1() bool {
    var v11 uint = 510
    v11 -= (v11 >> ((v11) & 63))
    for fi12 := 0; fi12 < 3; fi12++ {
        if false {
            for fi13 := 0; fi13 < 8; fi13++ {
                if true {
                    switch "a" {
                    case "hello":
                        v11 = 1
                    case "":
                        var v14 uint64 = 2
                        _ = v14
                    }
                    var v15 int64 = -421
                    _ = v15
                }
                v11 = v11
                _ = fi13
            }
        } else {
            v16 := []int{fi12}
            v11 -= v11
            _ = v16
        }
        _ = fi12
    }
    v17 := GenS2{F0: 32767}
    fmt.Println(v11)
    _ = v11
    _ = v17
    return false
}

func main() {
    var v1 bool = false
    var v2 int = 3
    v3 := make(map[string]string)
    v4 := make(map[uint32]string)
    v5 := []uint32{324, 4294967295, 42}
    v6 := make(map[uint32]string)
    if (!v1) {
        delete(v3, "foo bar")
        if (42 >= 643) {
            v2 = (v2 << ((2) & 63))
            v1 = false
        }
        fmt.Println(v2, v2)
    } else {
        switch v2 {
        case 0:
            var v18 string = "a"
            _ = v18
        }
    }
    if ((((v1 || (9223372036854775807 != 0)) || false) && (847 > 4294967295)) && (((-128 != -1) || false) && v1)) {
        switch (v2 << ((255) & 63)) {
        case 42:
            v1 = v1
        default:
            if true {
                var v19 string = ("foo bar" + ("hello" + "foo bar"))
                _ = v19
            }
        }
    }
    v1 = v1
    var v20 float64 = (3.14159 + ((((1e10 * (1e10 * -2.25)) - float64(float32(0.0))) * (1e10 - float64(-2.25))) + (-2.25 * 1e10)))
    v20 = float64((float32((v20 / 1.0)) / 3.14159))
    v21 := GenS1{F0: 0, F1: 1}
    fmt.Println(v2)
    v2 = (-(v2))
    for fi22 := 0; fi22 < 20; fi22++ {
        for fi23 := 0; fi23 < 5; fi23++ {
            v24 := GenS3{F0: 4294967295, F1: 96, F2: 2147483648}
            v3["hello"] = ""
            var v25 int32 = 1
            var v26 int32 = (v25 & (v25 - (v25 - v25)))
            _ = fi23
            _ = v24
            _ = v25
            _ = v26
        }
        for fi27 := 0; fi27 < 8; fi27++ {
            if (v1 && v1) {
                v28 := GenS3{F0: 42, F1: 0, F2: 2147483647}
                if (false && (3.14159 != v20)) {
                    switch "hello" {
                    case "hello":
                        v1 = (1 != 0)
                    }
                    v29 := GenS3{F0: 2147483647, F1: -65, F2: 50}
                    switch "hello" {
                    case "\n\t":
                        if (float64(((float32(v20) - 0.0) + ((0.1 / 1.0) - (3.14159 / -2.25)))) > (v20 / v20)) {
                            v2 = 0
                        }
                    case "hello":
                        var v30 uint = 502
                        _ = v30
                    }
                    _ = v29
                } else {
                    var v31 int16 = 32767
                    var v32 int = (fi22 & (^(fi22)))
                    _ = v31
                    _ = v32
                }
                _ = len(v3)
                _ = v28
            } else {
                if (-124 >= -128) {
                    v33 := GenS1{F0: 917, F1: 127}
                    _ = v33
                } else {
                    v5 = append(v5, 1)
                    v20 += (float64((0.0 + (0.1 - 3.14159))) / ((v20 + float64(3.14159)) + 3.14159))
                    v2 = fi22
                }
            }
            v5 = append(v5, 2147483647)
            _ = fi27
        }
        v5 = append(v5, 360)
        _ = fi22
    }
    v34 := GenS4{F0: 42, F1: 1}
    v35 := GenS1{F0: 1, F1: -128}
    var v36 float64 = v20
    for fi37 := 0; fi37 < 12; fi37++ {
        if len(v5) > 0 {
            v5[0] = 42
        }
        _ = fi37
    }
    if v1 {
        var v38 int16 = 32767
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
