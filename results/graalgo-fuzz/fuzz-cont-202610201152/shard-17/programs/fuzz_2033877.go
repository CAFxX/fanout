package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 uint64; F2 uint }
func genFn0() uint8 {
    for fi5 := 0; fi5 < 20; fi5++ {
        v6 := []int{(fi5 & -9223372036854775808), (fi5 ^ (fi5 >> ((658) & 63)))}
        if true {
            if (122 >= 606) {
                v7 := GenS1{F0: 0}
                _ = len(v6)
                var v8 int8 = -66
                _ = v7
                _ = v8
            }
            if len(v6) > 0 {
                v6[0] = fi5
            }
            fmt.Println(fi5)
        } else {
            if (46 < 0) {
                var v9 int16 = 32767
                _ = v9
            }
            var v10 uint64 = 451
            v6 = append(v6, (fi5 ^ (fi5 >> ((384) & 63))))
            _ = v10
        }
        _ = fi5
        _ = v6
    }
    var v11 bool = false
    _ = v11
    return 2
}

func main() {
    v1 := make(map[string]string)
    var v2 uint16 = 255
    var v3 int8 = 127
    v4 := []string{"", "hello"}
    for fi12 := 0; fi12 < 9; fi12++ {
        var v13 int = (fi12 * (fi12 % ((fi12) | 1)))
        var v14 int32 = 2147483647
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 uint = 4294967295
    var v16 uint = 18446744073709551615
    delete(v1, "")
    for fi17 := 0; fi17 < 2; fi17++ {
        if (216 != 2) {
            v18 := GenS2{F0: 42}
            var v19 string = ("\n\t" + "a")
            if len(v4) > 0 {
                v4[0] = v19
            }
            _ = v18
            _ = v19
        } else {
            v20 := GenS1{F0: -1}
            v2 = v2
            _ = v20
        }
        v16 = v15
        _ = fi17
    }
    if true {
        var v21 int = -9223372036854775808
        var v22 int32 = -2147483648
        _ = len(v4)
        _ = v21
        _ = v22
    } else {
        var v23 bool = ((128 < 0) && (-2147483648 == 128))
        v16 -= v15
        var v24 float64 = 3.14159
        _ = v23
        _ = v24
    }
    v25 := GenS3{F0: "\n\t", F1: 18446744073709551615, F2: 1}
    if ((!("" <= "hello")) && true) {
        if (true && (((3.14159 - (-2.25 - 3.14159)) * 3.14159) >= 1.5)) {
            var v26 uint32 = 206
            var v27 float32 = 1e10
            for fi28 := 0; fi28 < 20; fi28++ {
                for fi29 := 0; fi29 < 5; fi29++ {
                    v2 -= (v2 + v2)
                    _ = fi29
                }
                if (26 != v2) {
                    var v30 uint8 = uint8(v15)
                    _ = v30
                } else {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v31 float64 = (float64(v27) - -2.25)
                    var v32 int = fi28
                    _ = v31
                    _ = v32
                }
                _ = fi28
            }
            _ = v26
            _ = v27
        }
    } else {
        var v33 uint64 = 508
        var v34 string = "\n\t"
        v2 += 652
        _ = v33
        _ = v34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
