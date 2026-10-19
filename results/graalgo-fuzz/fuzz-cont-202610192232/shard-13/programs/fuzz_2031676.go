package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 int }
func genFn0(p0 int16) uint {
    if (!(!true)) {
        p0 += p0
        for fi6 := 0; fi6 < 1; fi6++ {
            v7 := []int{fi6, fi6, (fi6 - fi6), 128}
            _ = fi6
            _ = v7
        }
        if false {
            v8 := GenS1{F0: 0}
            _ = v8
        }
    }
    if true {
        v9 := []int{127}
        _ = v9
    }
    v10 := []int{}
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    _ = p0
    _ = v10
    return 42
}

func main() {
    var v1 float32 = 1.5
    v2 := []uint32{542, 2, 2147483647, 42}
    var v3 float64 = 0.0
    var v4 string = "a"
    v5 := make(map[string]uint32)
    var v11 int8 = 0
    for fi12 := 0; fi12 < 4; fi12++ {
        delete(v5, ("" + (v4 + v4)))
        switch v4 {
        case "\n\t":
            if (0.0 > v3) {
                if (!(v11 != -66)) {
                    if true {
                        if (!(!(941 <= 42))) {
                            v4 = v4
                        } else {
                            var v13 float64 = (-2.25 / v3)
                            v11 += (v11 >> ((2) & 63))
                            _ = v13
                        }
                        if true {
                            for fi14 := 0; fi14 < 19; fi14++ {
                                fmt.Println(v3, v4)
                                var v15 uint = 191
                                var v16 int64 = 42
                                v17 := GenS1{F0: v11}
                                _ = fi14
                                _ = v15
                                _ = v16
                                _ = v17
                            }
                            for fi18 := 0; fi18 < 1; fi18++ {
                                v19 := GenS1{F0: (v11 | (v11 | v11))}
                                _ = fi18
                                _ = v19
                            }
                        }
                    } else {
                        if true {
                            var v20 float32 = float32(float64((float32(v3) * float32(((v3 * float64(float32(3.14159))) - (v3 + v3))))))
                            _ = v20
                        } else {
                            if ((2147483647 <= 618) && ((v11 * int8(fi12)) == 127)) {
                                switch -77 {
                                case 32767:
                                    var v21 uint16 = 0
                                    _ = v21
                                case 0:
                                    v22 := v5[(v4 + v4)]
                                    _ = v22
                                case -9223372036854775808:
                                    v23 := GenS1{F0: 52}
                                    _ = v23
                                }
                                v11 += v11
                                fmt.Println(v1, v3, v11)
                            } else {
                                v24 := v5["hello"]
                                fmt.Println(v4)
                                _ = v24
                            }
                            var v25 string = v4
                            v11 = v11
                            _ = v25
                        }
                        v2 = append(v2, 716)
                    }
                }
                var v26 int = (fi12 << ((63) & 63))
                _ = v26
            }
        }
        var v27 uint = 1
        _ = fi12
        _ = v27
    }
    for fi28 := 0; fi28 < 10; fi28++ {
        var v29 bool = true
        var v30 int32 = 42
        v2 = append(v2, 65535)
        _ = fi28
        _ = v29
        _ = v30
    }
    v4 = v4
    if (1 >= 2) {
        fmt.Println(v3, v11)
        v11 += (v11 * v11)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        _ = len(v5)
        v5[("a" + v4)] = 42
    }
    fmt.Println(v3, v3)
    switch v4 {
    case "hello":
        if ("a" <= ("hello" + v4)) {
            v11 = 33
            if (!(1 == 18446744073709551615)) {
                var v31 uint = 555
                for fi32 := 0; fi32 < 2; fi32++ {
                    var v33 float64 = (((0.0 - 0.1) / v3) - v3)
                    _ = fi32
                    _ = v33
                }
                _ = v31
            } else {
                var v34 float64 = v3
                _ = v34
            }
        } else {
            var v35 int16 = 42
            _ = v35
        }
    case "a":
        switch "a" {
        case "\n\t":
            v36 := GenS1{F0: v11}
            _ = v36
        case "":
            v11 += 15
        case "foo bar":
            var v37 int = -2147483648
            _ = v37
        default:
            var v38 uint8 = 194
            _ = v38
        }
    }
    fmt.Println(v1)
    v39 := GenS2{F0: 42, F1: 715}
    if true {
        v40 := v5[(v4 + (("a" + v4) + "foo bar"))]
        _ = v40
    } else {
        for fi41 := 0; fi41 < 11; fi41++ {
            var v42 string = v4
            _ = fi41
            _ = v42
        }
        v4 = v4
    }
    switch ("\n\t" + "hello") {
    case "hello":
        if (-1 != 42) {
            fmt.Println(v4, v3, v4)
            v4 = v4
        }
    }
    var v43 float32 = v1
    if false {
        var v44 string = v4
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v43
    _ = v5
}
