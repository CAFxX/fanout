package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32 }

func main() {
    v1 := make(map[int]int)
    var v2 float64 = 3.14159
    var v3 int64 = -35
    v4 := make(map[string]uint32)
    var v5 uint16 = 2
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 uint = 0
        switch -1 {
        case -2147483648:
            for fi8 := 0; fi8 < 4; fi8++ {
                v9 := []int{-9223372036854775808, (fi6 & 473)}
                v10 := GenS1{F0: 0}
                _ = fi8
                _ = v10
                _ = v9
            }
        case 32768:
            if ((128 <= fi6) || (1e10 > (1.5 * (3.14159 + float32(v2))))) {
                v7 = (^(v7))
                if false {
                    var v11 uint32 = 262
                    _ = v11
                } else {
                    var v12 float32 = 0.1
                    for fi13 := 0; fi13 < 18; fi13++ {
                        if (uint8(v7) != 128) {
                            if true {
                                fmt.Println(fi13, v7)
                            }
                            v3 = (v3 << ((781) & 63))
                            v14 := GenS2{F0: -2147483648}
                            _ = v14
                        }
                        v15 := GenS1{F0: 1}
                        _ = fi13
                        _ = v15
                    }
                    v16 := []int{fi6, fi6}
                    _ = v12
                    _ = v16
                }
            } else {
                if (-2147483648 != 42) {
                    v17 := GenS2{F0: 1}
                    var v18 uint64 = 4294967295
                    if (false && false) {
                        for fi19 := 0; fi19 < 11; fi19++ {
                            v3 = v3
                            v20 := []int{int(v3), 1, 4294967295, int(v5)}
                            _ = fi19
                            _ = v20
                        }
                        fmt.Println(v2)
                    } else {
                        switch "hello" {
                        case "a":
                            switch 9223372036854775807 {
                            case -9223372036854775808:
                                var v21 bool = false
                                _ = v21
                            }
                        case "foo bar":
                            v1[(fi6 << (((v7 + (v7 + uint(v5)))) & 63))] = (fi6 >> (((v7 << ((v7) & 63))) & 63))
                        default:
                            switch (-(fi6)) {
                            case -1:
                                v3 += (v3 << ((512) & 63))
                            case 32768:
                                v22 := v1[127]
                                _ = v22
                            case -2147483648:
                                v7 = (v7 << ((uint(v18)) & 63))
                            default:
                                for fi23 := 0; fi23 < 9; fi23++ {
                                    v24 := v4["\n\t"]
                                    if (42 <= -1) {
                                        var v25 string = ("" + ("a" + ""))
                                        fmt.Println(v25, v25, fi23)
                                        _ = v25
                                    } else {
                                        if true {
                                            fmt.Println(v7, v5, v5)
                                            v24 = v24
                                        }
                                        var v26 float32 = float32((3.14159 + (v2 * 0.1)))
                                        _ = v26
                                    }
                                    _ = fi23
                                    _ = v24
                                }
                            }
                        }
                        v27 := GenS1{F0: -62}
                        if (!((!(32767 > (v5 % ((164) | 1)))) && true)) {
                            for fi28 := 0; fi28 < 20; fi28++ {
                                if false {
                                    v5 = (v5 * (v5 | v5))
                                }
                                _ = fi28
                            }
                            var v29 uint16 = v5
                            _ = v29
                        }
                        _ = v27
                    }
                    _ = v17
                    _ = v18
                }
            }
        case -233:
            for fi30 := 0; fi30 < 14; fi30++ {
                var v31 float32 = 0.0
                v2 -= ((0.0 / 1.0) - v2)
                var v32 uint8 = 232
                _ = fi30
                _ = v31
                _ = v32
            }
        default:
            var v33 bool = true
            _ = v33
        }
        _ = fi6
        _ = v7
    }
    for fi34 := 0; fi34 < 16; fi34++ {
        for fi35 := 0; fi35 < 10; fi35++ {
            var v36 uint16 = (^(v5))
            _ = fi35
            _ = v36
        }
        _ = fi34
    }
    var v37 uint32 = 72
    var v38 uint16 = (v5 | 42)
    for fi39 := 0; fi39 < 16; fi39++ {
        switch (("hello" + "\n\t") + "hello") {
        case "a":
            for fi40 := 0; fi40 < 12; fi40++ {
                fmt.Println(fi40, fi40)
                v38 -= 65535
                v2 = (-2.25 + v2)
                _ = fi40
            }
        case "foo bar":
            v41 := []int{}
            _ = v41
        case "":
            v38 += v38
        default:
            var v42 uint32 = v37
            _ = v42
        }
        var v43 int32 = -2147483648
        _ = fi39
        _ = v43
    }
    for fi44 := 0; fi44 < 20; fi44++ {
        var v45 float64 = float64(float32(v2))
        for fi46 := 0; fi46 < 18; fi46++ {
            v2 += v45
            _ = fi46
        }
        for fi47 := 0; fi47 < 2; fi47++ {
            v48 := []int{}
            _ = fi47
            _ = v48
        }
        _ = fi44
        _ = v45
    }
    v49 := GenS1{F0: 42}
    delete(v1, 0)
    var v50 float64 = ((v2 + ((v2 * -2.25) - (v2 + (((1e10 * v2) * float64(0.1)) - 0.0)))) * float64(1e10))
    fmt.Println(v50, v3, v38)
    v3 -= v3
    var v51 float64 = 0.1
    v38 = (v5 ^ (v38 / ((42) | 1)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
}
