package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 bool; F2 int64 }
type GenS3 struct { F0 int32; F1 int }
func genFn0(p0 uint64, p1 uint16, p2 uint64) int {
    var v5 string = "a"
    var v6 int16 = 32767
    if (true && false) {
        for fi7 := 0; fi7 < 20; fi7++ {
            v8 := GenS1{F0: 2}
            v9 := GenS2{F0: v5, F1: (((v5 + v5) + v5) >= v5), F2: 0}
            for fi10 := 0; fi10 < 9; fi10++ {
                var v11 string = v5
                _ = fi10
                _ = v11
            }
            _ = fi7
            _ = v8
            _ = v9
        }
        var v12 uint8 = 1
        fmt.Println(p2, p1)
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 683
}

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[string]string)
    v3 := make(map[int]int)
    var v4 bool = (!(false && false))
    v13 := GenS1{F0: 4294967295}
    var v14 int = -2147483648
    if (!false) {
        v14 = (v14 | v14)
        var v15 string = (("a" + "hello") + ("foo bar" + "a"))
        v4 = (1000 < 1)
        _ = v15
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        fmt.Println(fi16, v1)
        _ = fi16
    }
    var v17 uint8 = 0
    for fi18 := 0; fi18 < 19; fi18++ {
        var v19 uint32 = (^(v1))
        v19 += (v19 ^ (^(v1)))
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 1; fi20++ {
        var v21 uint = uint(v17)
        for fi22 := 0; fi22 < 4; fi22++ {
            if v4 {
                var v23 float64 = float64(1e10)
                _ = v23
            } else {
                if true {
                    var v24 uint8 = v17
                    v25 := []int{32767, v14, (v14 * genFn0(4294967295, 2, 370)), genFn0(517, 255, 1)}
                    _ = v24
                    _ = v25
                }
            }
            var v26 float64 = (-2.25 * float64(1.5))
            if v4 {
                if (1 != 2147483647) {
                    var v27 int32 = -522
                    v28 := v2[("\n\t" + "\n\t")]
                    var v29 bool = (0 <= 65)
                    var v30 int32 = (v27 + v27)
                    _ = v27
                    _ = v28
                    _ = v29
                    _ = v30
                }
                var v31 int8 = 42
                _ = v31
            }
            _ = fi22
            _ = v26
        }
        _ = fi20
        _ = v21
    }
    for fi32 := 0; fi32 < 9; fi32++ {
        for fi33 := 0; fi33 < 10; fi33++ {
            if (!true) {
                var v34 uint8 = v17
                fmt.Println(fi32, v34)
                _ = v34
            }
            _ = fi33
        }
        v35 := GenS1{F0: 4294967295}
        for fi36 := 0; fi36 < 14; fi36++ {
            _ = len(v3)
            if v4 {
                v1 -= v1
            } else {
                v37 := []int{genFn0(841, 65535, 1), (fi32 & (v14 >> ((uint(v17)) & 63))), genFn0(uint64(v1), 1, 1), (fi32 ^ (fi36 / (((fi36 + v14)) | 1)))}
                var v38 int16 = -32768
                _ = v37
                _ = v38
            }
            _ = fi36
        }
        _ = fi32
        _ = v35
    }
    var v39 bool = (!(((true || false) && (!((^(v1)) >= v1))) || (false || (104 <= 0))))
    v40 := GenS3{F0: -7, F1: v14}
    var v41 int = (v14 + (v14 * v14))
    var v42 uint64 = 2
    v43 := GenS1{F0: (^(v42))}
    v39 = v39
    delete(v2, "foo bar")
    var v44 uint = uint(v1)
    delete(v3, -135)
    var v45 float32 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
}
