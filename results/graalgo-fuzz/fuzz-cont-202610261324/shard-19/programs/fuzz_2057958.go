package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint64, p1 uint32, p2 uint) int {
    var v7 uint64 = p0
    var v8 uint = p2
    if ((3.14159 + 1.5) == 0.0) {
        v9 := []int{128, -2147483648, -328}
        var v10 float32 = (3.14159 * 0.0)
        v11 := GenS1{F0: (("\n\t" + "foo bar") < ("\n\t" + (("a" + "\n\t") + "\n\t")))}
        _ = v10
        _ = v11
        _ = v9
    } else {
        var v12 int32 = -1
        switch -1 {
        case 2147483647:
            for fi13 := 0; fi13 < 6; fi13++ {
                for fi14 := 0; fi14 < 2; fi14++ {
                    v15 := []int{(fi14 + 32768), fi14}
                    _ = fi14
                    _ = v15
                }
                var v16 int32 = v12
                var v17 uint32 = 0
                _ = fi13
                _ = v16
                _ = v17
            }
        }
        var v18 int64 = 821
        _ = v12
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 bool = (!(480 < 2))
    var v2 float32 = 1e10
    v3 := []uint32{424}
    v4 := make(map[uint32]string)
    v5 := make(map[uint32]bool)
    v6 := make(map[int]string)
    if false {
        for fi19 := 0; fi19 < 7; fi19++ {
            v20 := GenS2{F0: 6}
            switch "" {
            case "\n\t":
                v2 -= v2
            case "hello":
                fmt.Println(v1, v1)
            case "a":
                v1 = (-1 < 1)
            default:
                for fi21 := 0; fi21 < 11; fi21++ {
                    var v22 float64 = 1e10
                    _ = fi21
                    _ = v22
                }
            }
            var v23 float64 = 0.1
            _ = fi19
            _ = v20
            _ = v23
        }
        var v24 float64 = -2.25
        v24 += 0.1
        _ = v24
    }
    v1 = (-2.25 >= -2.25)
    for fi25 := 0; fi25 < 2; fi25++ {
        var v26 int32 = -2147483648
        var v27 int64 = 9223372036854775807
        switch "\n\t" {
        case "foo bar":
            _ = len(v6)
        }
        _ = fi25
        _ = v26
        _ = v27
    }
    for fi28 := 0; fi28 < 13; fi28++ {
        if (42 < 1) {
            for fi29 := 0; fi29 < 15; fi29++ {
                _ = len(v4)
                var v30 float64 = (3.14159 - 3.14159)
                var v31 uint8 = 1
                _ = fi29
                _ = v30
                _ = v31
            }
            if (false && (255 <= 0)) {
                v5[1] = (!(2 < 2))
                fmt.Println(v2, fi28, fi28)
                for fi32 := 0; fi32 < 16; fi32++ {
                    var v33 float64 = (0.0 + 1e10)
                    var v34 uint = 322
                    _ = fi32
                    _ = v33
                    _ = v34
                }
            } else {
                var v35 string = "foo bar"
                fmt.Println(v1)
                _ = v35
            }
        } else {
            var v36 uint16 = 255
            v37 := GenS1{F0: v1}
            var v38 int8 = -41
            _ = v36
            _ = v37
            _ = v38
        }
        if (722 <= 2147483647) {
            var v39 uint64 = 18446744073709551615
            _ = v39
        } else {
            v1 = v1
        }
        _ = fi28
    }
    fmt.Println(v1, v1)
    if ((true || (!(false || (v1 || false)))) && v1) {
        for fi40 := 0; fi40 < 3; fi40++ {
            switch ("foo bar" + "") {
            case "foo bar":
                fmt.Println(fi40)
            case "hello":
                var v41 uint16 = 490
                _ = v41
            case "":
                _ = len(v3)
            default:
                switch fi40 {
                case 128:
                    v2 -= -2.25
                case 4294967295:
                    v42 := GenS2{F0: 0}
                    _ = v42
                case 190:
                    switch (fi40 / (((fi40 % ((genFn0(506, 227, 214)) | 1))) | 1)) {
                    case 219:
                        if v1 {
                            _ = len(v3)
                            switch (fi40 % ((fi40) | 1)) {
                            case 223:
                                v2 += 0.0
                            default:
                                v1 = (("\n\t" + ("a" + "\n\t")) < "a")
                            }
                        }
                    }
                default:
                    for fi43 := 0; fi43 < 3; fi43++ {
                        fmt.Println(fi40, v1)
                        var v44 uint8 = 173
                        _ = fi43
                        _ = v44
                    }
                }
            }
            fmt.Println(v1)
            _ = fi40
        }
    }
    switch 128 {
    case -540:
        var v45 string = "foo bar"
        _ = v45
    default:
        var v46 bool = v1
        _ = v46
    }
    v2 -= 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
