package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 int32 }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = (!true)
    var v3 bool = (65535 == 4294967295)
    v4 := []string{}
    if false {
        if true {
            if v3 {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v5 string = "foo bar"
                _ = v5
            } else {
                v3 = (-914 >= -1)
                var v6 int = 1
                _ = v6
            }
            switch 190 {
            case -2147483648:
                _ = len(v1)
            case -9223372036854775808:
                var v7 int = 555
                _ = v7
            }
        } else {
            v3 = v2
            if len(v4) > 0 {
                v4[0] = ("hello" + "\n\t")
            }
        }
        if true {
            for fi8 := 0; fi8 < 3; fi8++ {
                if (!v2) {
                    if v2 {
                        delete(v1, "\n\t")
                        fmt.Println(v2, v2)
                        v3 = false
                    } else {
                        var v9 uint = 65535
                        var v10 float32 = float32(((1.5 - (((((((1.5 + 3.14159) * ((-2.25 - 1.5) * 0.0)) + float64(0.0)) / 0.1) / (-2.25 * 3.14159)) - (0.0 - (0.1 - 0.0))) - 0.0)) * 1e10))
                        _ = v10
                        _ = v9
                    }
                    v11 := GenS1{F0: 2147483647, F1: 394, F2: -722}
                    for fi12 := 0; fi12 < 2; fi12++ {
                        var v13 uint64 = 18446744073709551615
                        _ = fi12
                        _ = v13
                    }
                    _ = v11
                }
                if false {
                    v2 = (0.1 >= -2.25)
                }
                for fi14 := 0; fi14 < 9; fi14++ {
                    v15 := GenS1{F0: 743, F1: 2, F2: -1}
                    if (v2 && v3) {
                        var v16 bool = (!(153 < 18446744073709551615))
                        var v17 float64 = 3.14159
                        fmt.Println(v17, v16)
                        _ = v16
                        _ = v17
                    } else {
                        fmt.Println(fi14, v2)
                        var v18 int = (fi14 % (((fi14 ^ 766)) | 1))
                        _ = v18
                    }
                    _ = fi14
                    _ = v15
                }
                _ = fi8
            }
            v3 = v2
            var v19 uint = 543
            _ = v19
        } else {
            v2 = false
        }
    } else {
        v3 = v2
        var v20 float32 = (3.14159 - (3.14159 + ((float32(1e10) - 0.1) * 1e10)))
        var v21 int = 32767
        _ = v20
        _ = v21
    }
    v22 := GenS2{F0: 193}
    for fi23 := 0; fi23 < 19; fi23++ {
        fmt.Println(v3)
        var v24 int64 = 1
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 19; fi25++ {
        if (v3 && true) {
            var v26 int32 = 2147483647
            _ = v26
        }
        _ = fi25
    }
    var v27 uint32 = 2147483648
    v3 = false
    for fi28 := 0; fi28 < 4; fi28++ {
        v4 = append(v4, ("" + "\n\t"))
        var v29 int64 = 9223372036854775807
        delete(v1, ("hello" + "a"))
        _ = fi28
        _ = v29
    }
    var v30 int32 = 42
    fmt.Println(v30, v2)
    var v31 uint16 = 73
    v2 = v3
    v2 = (!v3)
    v3 = v2
    if v3 {
        var v32 uint8 = 0
        if (!v3) {
            var v33 uint = 1
            v34 := GenS1{F0: (v27 / ((v27) | 1)), F1: v31, F2: v30}
            _ = v33
            _ = v34
        }
        _ = v32
    } else {
        var v35 uint64 = 42
        if len(v4) > 0 {
            v4[0] = "a"
        }
        _ = v35
    }
    for fi36 := 0; fi36 < 6; fi36++ {
        v37 := GenS2{F0: 0}
        var v38 uint64 = 4294967295
        _ = fi36
        _ = v37
        _ = v38
    }
    v1["a"] = (v27 + v27)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
