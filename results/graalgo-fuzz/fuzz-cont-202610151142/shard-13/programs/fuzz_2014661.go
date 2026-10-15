package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint32 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 uint64; F1 uint64 }
func genFn0() string {
    switch -1 {
    case 875:
        var v7 int16 = 0
        _ = v7
    case -748:
        if (false && true) {
            var v8 uint16 = 65535
            if (false || (!(!false))) {
                v8 -= (v8 - v8)
            } else {
                v8 = 255
                if (1 <= 1) {
                    var v9 int16 = -32768
                    for fi10 := 0; fi10 < 11; fi10++ {
                        v11 := []int{(fi10 ^ fi10), fi10, (fi10 ^ (fi10 << ((617) & 63)))}
                        v11 = append(v11, (^(fi10)))
                        _ = fi10
                        _ = v11
                    }
                    v12 := GenS1{F0: 2, F1: 100, F2: 981}
                    _ = v12
                    _ = v9
                }
            }
            _ = v8
        } else {
            var v13 uint8 = 2
            v14 := GenS2{F0: 2, F1: v13}
            var v15 int64 = 0
            _ = v13
            _ = v14
            _ = v15
        }
    case 0:
        v16 := []int{32768, 414, 32768}
        _ = v16
    default:
        for fi17 := 0; fi17 < 8; fi17++ {
            for fi18 := 0; fi18 < 7; fi18++ {
                v19 := GenS2{F0: 236, F1: 1}
                _ = fi18
                _ = v19
            }
            switch "\n\t" {
            case "hello":
                v20 := []int{fi17}
                _ = v20
            case "a":
                if (0 < 0) {
                    var v21 int64 = -281
                    var v22 bool = true
                    var v23 bool = false
                    _ = v21
                    _ = v22
                    _ = v23
                }
            case "foo bar":
                if ((!("hello" > "\n\t")) && false) {
                    if false {
                        var v24 uint64 = 1
                        _ = v24
                    }
                } else {
                    var v25 bool = (0.0 <= 0.1)
                    _ = v25
                }
            }
            _ = fi17
        }
    }
    for fi26 := 0; fi26 < 3; fi26++ {
        v27 := []int{(fi26 - (fi26 << ((641) & 63))), 128}
        var v28 uint16 = 65535
        _ = fi26
        _ = v27
        _ = v28
    }
    var v29 uint8 = 1
    _ = v29
    return "foo bar"
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]int)
    v3 := make(map[int]uint32)
    var v4 int64 = 0
    var v5 int8 = 42
    v6 := []string{"a", ("a" + ("\n\t" + "foo bar")), (("\n\t" + "a") + "\n\t")}
    v30 := GenS3{F0: 385, F1: 328}
    v5 += v5
    var v31 float32 = 3.14159
    var v32 int = 32767
    for fi33 := 0; fi33 < 10; fi33++ {
        v5 = 0
        v34 := GenS1{F0: 224, F1: (v5 % ((v5) | 1)), F2: 0}
        var v35 bool = (-1 > 42)
        _ = fi33
        _ = v34
        _ = v35
    }
    v36 := v1[(^(v32))]
    v36 = (v36 << ((987) & 63))
    fmt.Println(v36, v5, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
