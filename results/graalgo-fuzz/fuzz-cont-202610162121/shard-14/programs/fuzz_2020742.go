package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]int)
    v3 := make(map[string]uint32)
    var v4 string = ""
    v4 = v4
    switch v4 {
    case "hello":
        if (1.5 == 1e10) {
            v4 = ("\n\t" + "foo bar")
        } else {
            if (453 > 42) {
                var v5 bool = (!(294 >= 236))
                for fi6 := 0; fi6 < 16; fi6++ {
                    v7 := GenS1{F0: 0}
                    var v8 uint32 = 423
                    var v9 uint8 = 151
                    _ = fi6
                    _ = v7
                    _ = v8
                    _ = v9
                }
                var v10 int32 = 0
                _ = v10
                _ = v5
            } else {
                var v11 uint8 = 42
                _ = v11
            }
            var v12 float64 = (1.5 + (1.5 * (float64((1e10 / 1.5)) + (1e10 * 1.5))))
            _ = v12
        }
    }
    var v13 string = "foo bar"
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 int64 = 0
        _ = fi14
        _ = v15
    }
    v13 = (v13 + v13)
    v13 = "\n\t"
    v13 = (v4 + v4)
    var v16 int32 = -1
    if true {
        v13 = "\n\t"
        switch ("\n\t" + v4) {
        case "a":
            v17 := GenS1{F0: 398}
            _ = v17
        }
    } else {
        switch 32768 {
        case -149:
            v18 := v2["foo bar"]
            _ = v18
        case 4294967295:
            v19 := GenS1{F0: 2}
            _ = v19
        }
    }
    v20 := []int{-9223372036854775808, 127}
    _ = len(v3)
    var v21 uint16 = 255
    if len(v20) >= 2 {
        v20 = v20[1:len(v20)]
    }
    switch 4294967295 {
    case -288:
        delete(v3, "foo bar")
    case -9223372036854775808:
        v13 = (v4 + (((v13 + "hello") + v13) + (((v13 + v13) + ((("\n\t" + (v4 + (v13 + (v13 + (("foo bar" + "hello") + "\n\t"))))) + v13) + v4)) + v4)))
    case 2147483647:
        v22 := GenS1{F0: 101}
        _ = v22
    }
    if true {
        switch 128 {
        case -2147483648:
            if false {
                v16 = 0
                switch "" {
                case "\n\t":
                    _ = len(v20)
                default:
                    fmt.Println(v4, v13)
                }
            }
        case 128:
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
        default:
            switch v4 {
            case "hello":
                for fi23 := 0; fi23 < 19; fi23++ {
                    var v24 uint8 = 42
                    v3[((v13 + ("a" + (v13 + v4))) + ((v4 + (v13 + (v4 + v4))) + "\n\t"))] = 4294967295
                    _ = fi23
                    _ = v24
                }
            }
        }
        if false {
            var v25 int = -9223372036854775808
            if false {
                v26 := GenS2{F0: false}
                v4 = "hello"
                _ = v26
            }
            if true {
                if true {
                    if len(v20) >= 2 {
                        v20 = v20[1:len(v20)]
                    }
                    var v27 uint64 = 4294967295
                    var v28 int16 = 42
                    _ = v27
                    _ = v28
                } else {
                    for fi29 := 0; fi29 < 17; fi29++ {
                        var v30 uint64 = 4294967295
                        _ = fi29
                        _ = v30
                    }
                    var v31 bool = (true || (!false))
                    var v32 bool = v31
                    _ = v31
                    _ = v32
                }
                for fi33 := 0; fi33 < 14; fi33++ {
                    for fi34 := 0; fi34 < 3; fi34++ {
                        var v35 float64 = 1.5
                        _ = fi34
                        _ = v35
                    }
                    v21 += (v21 + v21)
                    v16 -= 0
                    _ = fi33
                }
            } else {
                v36 := GenS2{F0: false}
                v37 := GenS1{F0: 42}
                _ = v36
                _ = v37
            }
            _ = v25
        } else {
            fmt.Println(v21, v13)
        }
        switch (v4 + "") {
        case "\n\t":
            v2[(v4 + ((("foo bar" + v13) + (v4 + (v13 + v4))) + ("" + v13)))] = 381
        case "a":
            _ = len(v1)
        case "":
            if true {
                if len(v20) > 0 {
                    v20[0] = 4294967295
                }
            }
        default:
            v4 = v13
        }
    } else {
        v13 = v13
        v4 = (v4 + "foo bar")
        var v38 float32 = float32(0.1)
        _ = v38
    }
    v39 := GenS1{F0: 4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v39
    _ = v4
}
