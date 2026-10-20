package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int32 }
type GenS2 struct { F0 uint32; F1 uint64; F2 int8 }
type GenS3 struct { F0 int; F1 uint32; F2 int64 }

func main() {
    v1 := make(map[int]string)
    var v2 float64 = float64(-2.25)
    var v3 uint = 42
    var v4 string = (("\n\t" + "hello") + "\n\t")
    var v5 bool = false
    for fi6 := 0; fi6 < 6; fi6++ {
        for fi7 := 0; fi7 < 14; fi7++ {
            var v8 int16 = -32768
            v3 = v3
            var v9 uint = (^(v3))
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    var v10 string = ((v4 + "\n\t") + (((((v4 + ((v4 + v4) + v4)) + (v4 + v4)) + "foo bar") + v4) + "\n\t"))
    v2 = 1.5
    v11 := GenS1{F0: 52, F1: v3, F2: 101}
    fmt.Println(v3, v3)
    switch v4 {
    case "foo bar":
        for fi12 := 0; fi12 < 7; fi12++ {
            v13 := []int{(fi12 * (fi12 / (((^(fi12))) | 1))), fi12, (fi12 + fi12), (^(fi12))}
            for fi14 := 0; fi14 < 16; fi14++ {
                v15 := GenS1{F0: 18446744073709551615, F1: 477, F2: -1}
                v16 := GenS1{F0: 4294967295, F1: (v3 ^ (v3 % ((v3) | 1))), F2: 1}
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = fi12
            _ = v13
        }
    case "a":
        if v5 {
            var v17 uint64 = 84
            _ = v17
        }
    default:
        v4 = (v4 + ((("hello" + v10) + v4) + v10))
    }
    if (-693 >= 32767) {
        if true {
            switch 0 {
            case 875:
                v1[42] = v4
            case 4294967295:
                var v18 int32 = -1
                _ = v18
            case 128:
                v19 := GenS2{F0: 42, F1: 99, F2: 127}
                _ = v19
            }
        }
        if (!v5) {
            v20 := []int{-2147483648, 32768}
            _ = v20
        } else {
            if false {
                for fi21 := 0; fi21 < 14; fi21++ {
                    v3 -= v3
                    _ = fi21
                }
            } else {
                var v22 uint16 = 279
                v23 := GenS3{F0: 32768, F1: 870, F2: 9223372036854775807}
                _ = v22
                _ = v23
            }
            var v24 string = v10
            _ = v24
        }
        var v25 uint32 = 2
        _ = v25
    } else {
        for fi26 := 0; fi26 < 15; fi26++ {
            var v27 int32 = -1
            delete(v1, fi26)
            var v28 int = (fi26 + (fi26 - 1))
            _ = fi26
            _ = v27
            _ = v28
        }
        var v29 int64 = 1
        _ = v29
    }
    var v30 uint64 = 260
    for fi31 := 0; fi31 < 1; fi31++ {
        if (v5 && (!(v3 == (v3 >> ((v3) & 63))))) {
            v5 = true
            for fi32 := 0; fi32 < 20; fi32++ {
                _ = len(v1)
                _ = fi32
            }
            v33 := GenS2{F0: 42, F1: (v30 ^ 4294967295), F2: 42}
            _ = v33
        } else {
            v34 := []int{fi31, fi31, 5, fi31}
            _ = v34
        }
        v35 := GenS1{F0: v30, F1: 255, F2: 28}
        _ = fi31
        _ = v35
    }
    var v36 uint32 = 2
    v1[9223372036854775807] = ""
    fmt.Println(v36)
    v37 := []int{32767}
    fmt.Println(v5, v4, v36)
    v30 -= 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v30
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
