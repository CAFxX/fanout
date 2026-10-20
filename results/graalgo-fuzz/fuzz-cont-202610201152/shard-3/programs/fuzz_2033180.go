package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 string; F2 uint8 }
func genFn0(p0 bool, p1 int16, p2 int8) uint8 {
    if p0 {
        p2 = -57
    }
    if (255 >= 0) {
        var v4 float64 = 3.14159
        _ = v4
    } else {
        for fi5 := 0; fi5 < 13; fi5++ {
            p0 = p0
            v6 := []int{}
            var v7 uint = 1
            _ = fi5
            _ = v6
            _ = v7
        }
        v8 := []int{-1, 222, 734}
        _ = v8
    }
    v9 := GenS1{F0: (32768 == -9223372036854775808), F1: (p1 * int16(p2)), F2: 586}
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 127
}
func genFn1() int {
    v10 := []int{955, -336, 32767, 128}
    var v11 int32 = 42
    if len(v10) > 0 {
        v10[0] = -56
    }
    for fi12 := 0; fi12 < 13; fi12++ {
        fmt.Println(v11)
        for fi13 := 0; fi13 < 3; fi13++ {
            v11 += v11
            switch (^(fi12)) {
            case 2147483647:
                v14 := GenS1{F0: ((!true) && true), F1: 1, F2: 0}
                _ = v14
            }
            _ = fi13
        }
        v11 -= v11
        _ = fi12
    }
    _ = v10
    _ = v11
    return -392
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]string)
    v3 := make(map[string]bool)
    var v15 uint8 = 2
    v16 := GenS2{F0: 4294967295}
    v15 = v15
    if false {
        for fi17 := 0; fi17 < 14; fi17++ {
            fmt.Println(v15, fi17)
            for fi18 := 0; fi18 < 3; fi18++ {
                v19 := GenS3{F0: v15}
                v15 += 243
                if false {
                    switch (-(fi17)) {
                    case 127:
                        var v20 int64 = -1
                        _ = v20
                    case 4294967295:
                        if false {
                            switch fi18 {
                            case 32767:
                                v21 := []int{(fi17 / ((fi17) | 1)), (fi18 + (fi17 << ((960) & 63))), -1, genFn1()}
                                _ = v21
                            case -525:
                                v22 := GenS1{F0: false, F1: 42, F2: 4294967295}
                                _ = v22
                            case 42:
                                v23 := []int{(^(fi18)), (fi18 & (fi18 + 42)), 181, fi17}
                                _ = v23
                            }
                            v15 -= (v15 / ((v15) | 1))
                            v15 += (^(v15))
                        } else {
                            v15 = (v15 + 127)
                        }
                    }
                    v24 := []int{genFn1()}
                    var v25 int32 = 0
                    var v26 uint = 0
                    var v27 float64 = ((0.0 / 1e10) / 1.5)
                    _ = v24
                    _ = v25
                    _ = v26
                    _ = v27
                } else {
                    if true {
                        var v28 uint = 292
                        v15 = v15
                        _ = v28
                    }
                    v29 := []int{fi17}
                    _ = v29
                }
                var v30 uint8 = v15
                _ = fi18
                _ = v19
                _ = v30
            }
            if false {
                fmt.Println(v15, v15)
                fmt.Println(v15)
            }
            _ = fi17
        }
    } else {
        v15 += 42
        for fi31 := 0; fi31 < 16; fi31++ {
            var v32 bool = (2 == 2)
            _ = fi31
            _ = v32
        }
    }
    var v33 uint64 = 0
    v3["a"] = (false || true)
    v33 = v33
    v33 += v33
    if (v15 == v15) {
        v33 -= v33
        for fi34 := 0; fi34 < 19; fi34++ {
            if false {
                v33 = (v33 & (v33 >> ((979) & 63)))
            } else {
                v35 := []int{fi34}
                v2["a"] = ((("" + ("hello" + "\n\t")) + "hello") + "a")
                _ = v35
            }
            var v36 int = (fi34 >> ((4294967295) & 63))
            _ = fi34
            _ = v36
        }
        var v37 int16 = -256
        _ = v37
    } else {
        v3[""] = (!((0.0 / 0.1) == ((3.14159 * 1e10) * -2.25)))
    }
    v38 := GenS4{F0: 223, F1: "a", F2: v15}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v33
    _ = v38
}
