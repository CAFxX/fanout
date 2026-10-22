package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 int64; F1 uint16 }
func genFn0(p0 int8, p1 int16, p2 bool) uint16 {
    p0 += (p0 | (p0 & (p0 << ((1) & 63))))
    v7 := []int{128, 381, 794}
    if true {
        switch 2147483647 {
        case 1:
            var v8 uint64 = 557
            _ = v8
        default:
            fmt.Println(p2)
        }
        if (1e10 < 1e10) {
            fmt.Println(p1)
        } else {
            fmt.Println(p2, p0)
        }
    } else {
        if (p2 || (755 != 136)) {
            var v9 float32 = float32(0.1)
            if (!(float32(float64(v9)) >= (float32(0.0) + v9))) {
                if p2 {
                    v10 := GenS1{F0: 255, F1: 2147483647, F2: -109}
                    if (127 < 106) {
                        for fi11 := 0; fi11 < 13; fi11++ {
                            var v12 uint64 = 2
                            _ = fi11
                            _ = v12
                        }
                        for fi13 := 0; fi13 < 19; fi13++ {
                            v14 := GenS1{F0: 91, F1: 632, F2: fi13}
                            p2 = true
                            _ = fi13
                            _ = v14
                        }
                    }
                    var v15 float32 = v9
                    _ = v10
                    _ = v15
                }
                v16 := GenS1{F0: 65, F1: 2, F2: 32768}
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                _ = v16
            }
            _ = v9
        }
        var v17 uint8 = 42
        _ = v17
    }
    for fi18 := 0; fi18 < 9; fi18++ {
        if true {
            p2 = p2
        }
        _ = fi18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}

func main() {
    var v1 int8 = 5
    v2 := []string{"\n\t", "a", ("hello" + ("" + "\n\t"))}
    var v3 int16 = 42
    v4 := make(map[uint32]uint32)
    v5 := make(map[uint32]int)
    var v6 int16 = v3
    v19 := GenS2{F0: 4294967295, F1: 128}
    var v20 int16 = (v6 >> ((255) & 63))
    switch "" {
    case "a":
        var v21 string = ("\n\t" + "")
        _ = v21
    }
    if false {
        fmt.Println(v20, v3)
        if true {
            for fi22 := 0; fi22 < 11; fi22++ {
                if false {
                    switch -1 {
                    case -1:
                        var v23 uint8 = 255
                        _ = v23
                    }
                    var v24 string = "\n\t"
                    if ((812 < 1) && (!(!(!(!true))))) {
                        v1 += (-(v1))
                    }
                    _ = v24
                }
                for fi25 := 0; fi25 < 1; fi25++ {
                    fmt.Println(fi22, v1, v3)
                    v26 := GenS3{F0: (v1 + v1), F1: 42}
                    _ = fi25
                    _ = v26
                }
                v27 := GenS4{F0: -9223372036854775808, F1: 65535}
                _ = fi22
                _ = v27
            }
            var v28 int32 = 42
            v5[568] = 32767
            _ = v28
        } else {
            var v29 int64 = -1
            _ = v29
        }
        fmt.Println(v6, v3, v20)
    } else {
        v30 := GenS1{F0: 0, F1: 33, F2: -1}
        _ = v30
    }
    var v31 float32 = 3.14159
    var v32 string = "foo bar"
    if (2 > 2) {
        fmt.Println(v3, v3, v20)
    } else {
        if false {
            v3 = (v3 & 32767)
        } else {
            fmt.Println(v31)
            if true {
                var v33 uint8 = 128
                delete(v4, 2147483647)
                _ = v33
            } else {
                for fi34 := 0; fi34 < 8; fi34++ {
                    v20 = v20
                    _ = fi34
                }
            }
            var v35 int32 = 42
            _ = v35
        }
        var v36 int32 = 266
        fmt.Println(v32)
        _ = v36
    }
    v3 = (v20 & v3)
    v37 := v5[65535]
    var v38 uint16 = genFn0(v1, (v20 + 1), true)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v31
    _ = v32
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
