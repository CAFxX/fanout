package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 int32; F2 uint32 }
type GenS3 struct { F0 uint32; F1 uint8; F2 string }
func genFn0(p0 uint64, p1 int16, p2 string) uint32 {
    v7 := GenS1{F0: 42}
    var v8 int32 = 0
    var v9 uint = 255
    var v10 string = "foo bar"
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 2
}
func genFn1() uint {
    v11 := GenS1{F0: -678}
    v12 := GenS2{F0: 0, F1: -2147483648, F2: 168}
    v13 := GenS2{F0: -1, F1: 0, F2: 4294967295}
    var v14 uint32 = 4294967295
    v14 = v14
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    return 255
}

func main() {
    v1 := []int{}
    var v2 bool = (!(2 >= 1))
    var v3 float32 = 1e10
    v4 := make(map[string]uint32)
    var v5 bool = v2
    v6 := []string{"hello", "", ("\n\t" + ""), "a"}
    var v15 int = 64
    v16 := v4[""]
    v17 := GenS3{F0: (v16 >> ((1) & 63)), F1: 255, F2: "a"}
    v5 = true
    v1 = append(v1, v15)
    if true {
        delete(v4, ((("\n\t" + "a") + "foo bar") + "hello"))
    } else {
        switch 127 {
        case 0:
            for fi18 := 0; fi18 < 14; fi18++ {
                var v19 int16 = -532
                v19 -= (v19 + (v19 * (^(v19))))
                _ = fi18
                _ = v19
            }
        }
        if (false && ("foo bar" == ("\n\t" + ("\n\t" + "\n\t")))) {
            v6 = append(v6, "a")
        }
    }
    var v20 float64 = 1e10
    var v21 float64 = v20
    v5 = ((!v5) && (v5 || (!true)))
    v2 = v5
    for fi22 := 0; fi22 < 1; fi22++ {
        if (632 != 354) {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        }
        _ = fi22
    }
    switch "" {
    case "foo bar":
        for fi23 := 0; fi23 < 15; fi23++ {
            var v24 string = ""
            var v25 uint8 = 147
            var v26 bool = v2
            _ = fi23
            _ = v24
            _ = v25
            _ = v26
        }
    case "a":
        v16 = (^(v16))
    }
    var v27 uint8 = 0
    var v28 uint16 = 1
    for fi29 := 0; fi29 < 12; fi29++ {
        var v30 int = (fi29 % ((32767) | 1))
        _ = len(v4)
        _ = fi29
        _ = v30
    }
    v16 = genFn0(779, int16(v15), "a")
    v28 += v28
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
