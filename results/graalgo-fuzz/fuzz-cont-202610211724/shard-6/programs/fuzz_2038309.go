package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint; F1 uint; F2 int64 }
type GenS4 struct { F0 string }
func genFn0(p0 int, p1 int) int8 {
    v4 := []int{-1, 4294967295}
    for fi5 := 0; fi5 < 1; fi5++ {
        for fi6 := 0; fi6 < 14; fi6++ {
            var v7 int = (^(p1))
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v4
    return -1
}
func genFn1(p0 uint32, p1 uint) uint16 {
    p1 = 18446744073709551615
    if (((p1 - (p1 | (^(p1)))) != p1) && true) {
        switch "a" {
        case "a":
            var v8 uint64 = 18446744073709551615
            _ = v8
        }
        var v9 uint8 = 2
        switch ((("hello" + (("a" + (("\n\t" + ("hello" + "foo bar")) + ("hello" + "foo bar"))) + (("foo bar" + ("a" + ("foo bar" + "\n\t"))) + ("hello" + ("\n\t" + "a"))))) + "\n\t") + "a") {
        case "a":
            for fi10 := 0; fi10 < 8; fi10++ {
                for fi11 := 0; fi11 < 2; fi11++ {
                    p0 = p0
                    _ = fi11
                }
                for fi12 := 0; fi12 < 19; fi12++ {
                    v13 := []int{(fi12 + fi10), -2147483648}
                    _ = fi12
                    _ = v13
                }
                p1 = uint(p0)
                _ = fi10
            }
        case "\n\t":
            v14 := GenS1{F0: 100}
            _ = v14
        case "":
            v15 := GenS1{F0: -110}
            _ = v15
        default:
            var v16 uint8 = 1
            _ = v16
        }
        _ = v9
    }
    var v17 float32 = (float32((1e10 + (-2.25 - (3.14159 - 0.0)))) - 3.14159)
    fmt.Println(v17, v17, v17)
    _ = p0
    _ = p1
    _ = v17
    return 1
}

func main() {
    var v1 uint32 = 4294967295
    var v2 int16 = 1000
    v3 := make(map[int]int)
    var v18 uint16 = 0
    var v19 uint64 = 1
    v18 = genFn1(uint32(v19), 1)
    if (42 <= -267) {
        v20 := GenS1{F0: 57}
        _ = v20
    }
    v2 = -1
    v21 := GenS1{F0: 83}
    for fi22 := 0; fi22 < 1; fi22++ {
        if (!true) {
            fmt.Println(v2, v18)
            for fi23 := 0; fi23 < 3; fi23++ {
                v1 = v1
                v24 := GenS2{F0: 32767}
                var v25 float32 = -2.25
                _ = fi23
                _ = v24
                _ = v25
            }
        }
        for fi26 := 0; fi26 < 9; fi26++ {
            var v27 int32 = -381
            if false {
                if (!false) {
                    var v28 uint16 = 519
                    _ = v28
                } else {
                    v27 -= (v27 - (v27 & v27))
                    var v29 int8 = genFn0(fi26, fi22)
                    _ = v29
                }
                if true {
                    v30 := v3[fi22]
                    var v31 float64 = 0.1
                    v32 := GenS3{F0: 4294967295, F1: 822, F2: 9223372036854775807}
                    _ = v30
                    _ = v31
                    _ = v32
                }
            } else {
                v33 := []int{fi22}
                v19 = 0
                _ = v33
            }
            v34 := GenS4{F0: ""}
            _ = fi26
            _ = v27
            _ = v34
        }
        _ = fi22
    }
    v35 := GenS2{F0: (^(v18))}
    v3[186] = -453
    v19 = v19
    var v36 int32 = -922
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v35
    _ = v36
}
