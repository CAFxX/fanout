package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 int32; F2 int }

func main() {
    var v1 float32 = (0.0 - -2.25)
    var v2 int64 = 42
    v3 := []int{4294967295}
    switch (("a" + "hello") + (("hello" + "a") + (("hello" + "\n\t") + "hello"))) {
    case "foo bar":
        for fi4 := 0; fi4 < 18; fi4++ {
            if true {
                _ = len(v3)
            } else {
                var v5 uint64 = 762
                if (!false) {
                    v5 = 42
                    for fi6 := 0; fi6 < 4; fi6++ {
                        var v7 bool = true
                        _ = fi6
                        _ = v7
                    }
                    v8 := GenS1{F0: 19}
                    _ = v8
                }
                if true {
                    var v9 int64 = v2
                    _ = v9
                } else {
                    fmt.Println(fi4, v2, v1)
                }
                _ = v5
            }
            v2 = v2
            v10 := GenS1{F0: 128}
            _ = fi4
            _ = v10
        }
    case "":
        for fi11 := 0; fi11 < 14; fi11++ {
            var v12 uint = 18446744073709551615
            v1 -= v1
            _ = fi11
            _ = v12
        }
    }
    v13 := GenS1{F0: 2}
    if (false || true) {
        switch 4294967295 {
        case 11:
            if false {
                v2 = v2
                fmt.Println(v1, v2)
            } else {
                if len(v3) > 0 {
                    v3[0] = 42
                }
            }
        case -183:
            var v14 uint16 = 255
            _ = v14
        }
    } else {
        var v15 int = 128
        var v16 uint16 = 32767
        _ = len(v3)
        _ = v15
        _ = v16
    }
    var v17 uint64 = 4294967295
    v17 = (v17 ^ v17)
    v2 -= v2
    v1 = (v1 - (float32(1.5) + 0.1))
    if (42 <= 0) {
        var v18 uint = 2
        _ = len(v3)
        _ = v18
    } else {
        if false {
            if ("\n\t" == "a") {
                for fi19 := 0; fi19 < 2; fi19++ {
                    if (1e10 >= 1e10) {
                        var v20 int16 = -1
                        _ = v20
                    } else {
                        v17 = v17
                        v3 = append(v3, -192)
                        v2 += v2
                    }
                    _ = fi19
                }
                var v21 uint16 = 733
                _ = v21
            }
            var v22 uint = 399
            _ = v22
        } else {
            var v23 string = "a"
            _ = v23
        }
        v17 = (v17 + (v17 | (v17 % ((v17) | 1))))
        var v24 uint64 = v17
        _ = v24
    }
    fmt.Println(v17, v17)
    var v25 int32 = 0
    if (2 > 1) {
        v2 = int64(v17)
        var v26 int = 4294967295
        var v27 int32 = 475
        _ = v26
        _ = v27
    } else {
        for fi28 := 0; fi28 < 14; fi28++ {
            var v29 uint = 607
            var v30 uint = (v29 - 724)
            v1 += v1
            _ = fi28
            _ = v29
            _ = v30
        }
        if true {
            if false {
                for fi31 := 0; fi31 < 6; fi31++ {
                    if (v17 == v17) {
                        var v32 uint8 = 2
                        var v33 uint16 = 1
                        _ = v32
                        _ = v33
                    }
                    _ = fi31
                }
                switch "foo bar" {
                case "foo bar":
                    for fi34 := 0; fi34 < 10; fi34++ {
                        fmt.Println(fi34, v17, v1)
                        _ = fi34
                    }
                case "hello":
                    v35 := GenS2{F0: (v17 << ((1) & 63)), F1: (v25 << ((297) & 63)), F2: 4294967295}
                    _ = v35
                default:
                    fmt.Println(v17)
                }
            } else {
                var v36 uint64 = 0
                v17 += (v17 >> ((42) & 63))
                _ = v36
            }
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v37 uint64 = (v17 % (((^(v17))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v25
    _ = v3
    _ = v37
}
