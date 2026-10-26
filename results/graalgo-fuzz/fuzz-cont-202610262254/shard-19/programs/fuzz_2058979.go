package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 bool, p1 uint64) int {
    var v5 int32 = -228
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := []int{(fi6 / (((fi6 + fi6)) | 1)), (fi6 / ((119) | 1))}
        _ = fi6
        _ = v7
    }
    if (1.5 == float32(1.5)) {
        p0 = true
    }
    if (((42 >= -998) && p0) && (!false)) {
        var v8 uint32 = 470
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v5
    return 32768
}

func main() {
    var v1 int8 = 0
    v2 := []string{("\n\t" + ""), "\n\t", "foo bar", "hello"}
    v3 := make(map[int]int)
    var v4 uint64 = 1
    var v9 float32 = 0.0
    v2 = append(v2, "foo bar")
    var v10 int64 = 612
    switch 816 {
    case 32768:
        if len(v2) > 0 {
            v2[0] = ""
        }
    case 0:
        var v11 uint8 = 128
        _ = v11
    }
    v12 := GenS1{F0: v4}
    if (!false) {
        for fi13 := 0; fi13 < 2; fi13++ {
            v1 = (v1 + v1)
            for fi14 := 0; fi14 < 14; fi14++ {
                var v15 int = (fi14 * (^(fi13)))
                var v16 float32 = 3.14159
                _ = fi14
                _ = v15
                _ = v16
            }
            var v17 bool = ((0.1 > 3.14159) && ((!(!true)) && (true || (!true))))
            _ = fi13
            _ = v17
        }
        switch "\n\t" {
        case "a":
            if (false && (18446744073709551615 >= 18446744073709551615)) {
                if ((v10 < (v10 & (v10 ^ v10))) && false) {
                    var v18 uint32 = 2147483648
                    switch "a" {
                    case "foo bar":
                        fmt.Println(v4)
                    }
                    if true {
                        if len(v2) > 0 {
                            v2[0] = (("foo bar" + "\n\t") + "a")
                        }
                        v19 := GenS1{F0: 33}
                        fmt.Println(v10)
                        _ = v19
                    }
                    _ = v18
                } else {
                    fmt.Println(v4, v1)
                    var v20 uint16 = 255
                    _ = v20
                }
                var v21 uint64 = 37
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v21
            } else {
                v10 += 9223372036854775807
                v22 := GenS1{F0: (v4 ^ v4)}
                var v23 int32 = 2147483647
                _ = v22
                _ = v23
            }
        case "\n\t":
            v4 = v4
        }
        v24 := GenS1{F0: (v4 + v4)}
        _ = v24
    }
    fmt.Println(v1, v9)
    if false {
        fmt.Println(v4)
        var v25 uint8 = 128
        _ = v25
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    fmt.Println(v4)
    fmt.Println(v1, v9, v4)
    _ = len(v2)
    delete(v3, -1)
    if false {
        for fi26 := 0; fi26 < 14; fi26++ {
            v27 := GenS1{F0: (v4 * v4)}
            _ = fi26
            _ = v27
        }
        v4 -= v4
    } else {
        if true {
            var v28 int = -150
            _ = v28
        }
        v10 = 0
    }
    switch 1 {
    case 1:
        for fi29 := 0; fi29 < 4; fi29++ {
            if (255 < 0) {
                var v30 int8 = (v1 | 100)
                var v31 uint = 4294967295
                v32 := GenS1{F0: 18446744073709551615}
                _ = v30
                _ = v31
                _ = v32
            } else {
                if false {
                    if (1 <= 1) {
                        fmt.Println(v1, fi29)
                        fmt.Println(v1)
                    }
                    var v33 float32 = 3.14159
                    fmt.Println(fi29, fi29)
                    _ = v33
                } else {
                    var v34 string = "foo bar"
                    _ = v34
                }
                for fi35 := 0; fi35 < 2; fi35++ {
                    v36 := GenS1{F0: 1}
                    _ = fi35
                    _ = v36
                }
            }
            _ = fi29
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
