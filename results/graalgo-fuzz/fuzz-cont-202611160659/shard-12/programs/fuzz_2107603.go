package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint32; F1 int64 }
func genFn0() int8 {
    if true {
        for fi5 := 0; fi5 < 10; fi5++ {
            v6 := []int{}
            _ = fi5
            _ = v6
        }
    }
    var v7 uint64 = 4294967295
    var v8 int32 = -819
    _ = v7
    _ = v8
    return -1
}
func genFn1(p0 bool) uint {
    for fi9 := 0; fi9 < 20; fi9++ {
        p0 = p0
        p0 = p0
        fmt.Println(p0)
        _ = fi9
    }
    var v10 int16 = -1
    var v11 int32 = -189
    if p0 {
        v12 := []int{925, -115, 905, 32767}
        _ = v12
    } else {
        var v13 string = "foo bar"
        _ = v13
    }
    v11 -= (v11 + v11)
    _ = p0
    _ = v10
    _ = v11
    return 255
}

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = true
    v3 := []string{"\n\t", ("foo bar" + "")}
    var v4 float64 = 1.5
    var v14 bool = v2
    fmt.Println(v14)
    v15 := GenS1{F0: -553, F1: 127}
    var v16 uint8 = 127
    var v17 int32 = 0
    v4 = (0.0 / ((1.5 / -2.25) - 1.5))
    if (953 <= 2) {
        var v18 uint8 = v16
        for fi19 := 0; fi19 < 10; fi19++ {
            for fi20 := 0; fi20 < 17; fi20++ {
                if (!v14) {
                    v21 := GenS1{F0: 1000, F1: 42}
                    _ = v21
                }
                if (float64(0.0) != v4) {
                    v22 := GenS2{F0: 4294967295, F1: 729}
                    _ = v22
                }
                if (1 == -94) {
                    v14 = ((719 >= 679) || (int16(v17) <= 436))
                }
                _ = fi20
            }
            _ = fi19
        }
        _ = v18
    }
    v4 += 1e10
    for fi23 := 0; fi23 < 2; fi23++ {
        if true {
            for fi24 := 0; fi24 < 8; fi24++ {
                v4 = ((v4 * v4) / 1.0)
                _ = fi24
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if (v2 || (v17 != v17)) {
                var v25 int8 = -1
                _ = v25
            }
        } else {
            var v26 uint = 2
            var v27 float64 = 0.0
            var v28 bool = (0 != (^(v17)))
            _ = v26
            _ = v27
            _ = v28
        }
        v2 = (true && (0.1 > ((1.5 / (0.0 - 0.1)) - ((float32((v4 * 3.14159)) + 0.0) + float32(v4)))))
        for fi29 := 0; fi29 < 13; fi29++ {
            if (v2 || v14) {
                var v30 int32 = (^(v17))
                var v31 int = -540
                v32 := GenS1{F0: int16(fi23), F1: 127}
                _ = v30
                _ = v31
                _ = v32
            } else {
                if ((v17 ^ (v17 - -2147483648)) > 0) {
                    var v33 bool = v2
                    v34 := GenS1{F0: 32767, F1: genFn0()}
                    _ = v33
                    _ = v34
                }
            }
            _ = fi29
        }
        _ = fi23
    }
    var v35 bool = (!(0 != 1))
    var v36 uint8 = 255
    v37 := GenS1{F0: 385, F1: 100}
    _ = len(v1)
    var v38 string = ""
    var v39 bool = false
    var v40 uint64 = 354
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
}
