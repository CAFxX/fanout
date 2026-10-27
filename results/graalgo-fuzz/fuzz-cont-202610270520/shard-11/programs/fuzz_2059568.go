package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 int16; F2 int8 }
type GenS3 struct { F0 uint64; F1 uint8; F2 int64 }
type GenS4 struct { F0 int32; F1 uint64; F2 uint }
func genFn0(p0 int, p1 uint32) int16 {
    if false {
        if ((p1 ^ (p1 % ((p1) | 1))) != p1) {
            v5 := GenS1{F0: 2}
            v6 := GenS1{F0: 686}
            v7 := GenS1{F0: 963}
            _ = v5
            _ = v6
            _ = v7
        }
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        if true {
            v9 := []int{849, 0, -2147483648}
            p0 = p0
            _ = v9
        } else {
            var v10 uint16 = 255
            v10 -= (v10 + 255)
            _ = v10
        }
        for fi11 := 0; fi11 < 15; fi11++ {
            var v12 int16 = 896
            fmt.Println(p0, p1)
            _ = fi11
            _ = v12
        }
        _ = fi8
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        v14 := GenS2{F0: 128, F1: 42, F2: 100}
        v15 := GenS3{F0: 526, F1: 42, F2: -103}
        _ = fi13
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        fmt.Println(p1, p1)
        var v17 uint64 = 18446744073709551615
        var v18 uint32 = uint32(v17)
        _ = fi16
        _ = v17
        _ = v18
    }
    _ = p0
    _ = p1
    return -1
}
func genFn1(p0 int8) int32 {
    var v19 bool = (("foo bar" + (("a" + "a") + ("\n\t" + ("foo bar" + ("\n\t" + "foo bar"))))) <= (("\n\t" + "") + "hello"))
    v19 = v19
    var v20 int16 = -874
    for fi21 := 0; fi21 < 3; fi21++ {
        v19 = v19
        switch ("foo bar" + "hello") {
        case "a":
            v22 := []int{fi21, (^(fi21)), 32768, (fi21 + fi21)}
            _ = v22
        case "foo bar":
            for fi23 := 0; fi23 < 19; fi23++ {
                if ((42 >= 2) || (v19 || (0 > 2147483647))) {
                    if v19 {
                        fmt.Println(fi21, p0)
                        var v24 int16 = 42
                        _ = v24
                    } else {
                        v25 := GenS1{F0: 682}
                        v26 := GenS3{F0: 359, F1: 127, F2: 377}
                        v19 = v19
                        v27 := GenS2{F0: 127, F1: genFn0((^(fi23)), 789), F2: p0}
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                    switch (fi23 + fi23) {
                    case 193:
                        fmt.Println(fi21, v20)
                    case 42:
                        v19 = (!v19)
                    }
                    var v28 bool = (("\n\t" + "foo bar") >= "foo bar")
                    _ = v28
                }
                _ = fi23
            }
        case "":
            v19 = v19
        }
        _ = fi21
    }
    v29 := []int{}
    _ = p0
    _ = v19
    _ = v20
    _ = v29
    return 0
}

func main() {
    var v1 int32 = 42
    v2 := []uint32{0}
    v3 := make(map[uint32]int)
    var v4 float32 = -2.25
    if ((v4 / float32(((float64(v4) / 1.0) * (0.1 + 0.1)))) <= ((0.1 * -2.25) - (v4 - float32((3.14159 * (3.14159 / 1.5)))))) {
        _ = len(v3)
        var v30 uint16 = 42
        _ = v30
    }
    v31 := GenS2{F0: 1, F1: 32767, F2: 127}
    v4 = 3.14159
    if true {
        for fi32 := 0; fi32 < 10; fi32++ {
            v1 = 1
            _ = fi32
        }
    } else {
        v1 = -960
    }
    var v33 uint32 = 42
    var v34 string = ("\n\t" + "hello")
    var v35 bool = (((v34 + ((v34 + "a") + ((("hello" + v34) + v34) + v34))) != ("foo bar" + "hello")) || true)
    if (true && (!((v4 + (v4 - v4)) != 1e10))) {
        v4 = v4
    } else {
        var v36 uint64 = uint64(v33)
        v37 := GenS4{F0: (v1 ^ genFn1(92)), F1: 588, F2: 65535}
        for fi38 := 0; fi38 < 4; fi38++ {
            fmt.Println(v1)
            _ = fi38
        }
        _ = v36
        _ = v37
    }
    v33 += v33
    switch 9223372036854775807 {
    case -1:
        for fi39 := 0; fi39 < 20; fi39++ {
            delete(v3, (v33 + v33))
            _ = len(v2)
            var v40 string = v34
            _ = fi39
            _ = v40
        }
    }
    if (-248 <= -228) {
        v35 = false
        v1 = (v1 * (v1 + v1))
    } else {
        var v41 int16 = 42
        _ = v41
    }
    v35 = true
    var v42 int64 = 354
    switch -535 {
    case 1:
        var v43 float64 = ((0.0 - float64(0.0)) + ((float64(v4) + float64((v4 - (-2.25 / float32((1.5 / ((-2.25 - 3.14159) * 3.14159))))))) + 0.0))
        _ = v43
    }
    v2 = append(v2, 4294967295)
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v42
}
