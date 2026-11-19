package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 int16 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]int)
    var v3 float32 = (0.1 / 1.0)
    v4 := []int{128, 2147483647}
    v5 := GenS1{F0: 32767}
    fmt.Println(v3)
    v4 = append(v4, 2147483647)
    fmt.Println(v3)
    if true {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        if true {
            if (!false) {
                var v6 int32 = 42
                var v7 int32 = (v6 >> ((2) & 63))
                if (42 <= 41) {
                    var v8 uint = 255
                    if false {
                        for fi9 := 0; fi9 < 8; fi9++ {
                            var v10 int32 = (v6 * (v7 >> ((v8) & 63)))
                            v7 -= (^(v7))
                            _ = fi9
                            _ = v10
                        }
                        if true {
                            var v11 uint = 4294967295
                            v12 := GenS1{F0: 42}
                            _ = v11
                            _ = v12
                        }
                    } else {
                        v2[0] = 9223372036854775807
                    }
                    _ = v8
                } else {
                    for fi13 := 0; fi13 < 8; fi13++ {
                        switch 127 {
                        case 42:
                            for fi14 := 0; fi14 < 1; fi14++ {
                                var v15 uint = 65535
                                var v16 uint64 = 933
                                _ = fi14
                                _ = v15
                                _ = v16
                            }
                        case -26:
                            var v17 uint32 = 42
                            _ = v17
                        }
                        for fi18 := 0; fi18 < 10; fi18++ {
                            var v19 uint64 = 18446744073709551615
                            _ = fi18
                            _ = v19
                        }
                        _ = fi13
                    }
                }
                _ = v6
                _ = v7
            } else {
                v20 := GenS1{F0: 32767}
                for fi21 := 0; fi21 < 8; fi21++ {
                    var v22 int64 = -1
                    switch (fi21 - -931) {
                    case 4294967295:
                        fmt.Println(fi21)
                    case 42:
                        fmt.Println(fi21)
                    default:
                        v22 = v22
                    }
                    _ = fi21
                    _ = v22
                }
                switch ((("a" + ("foo bar" + ((("" + (((((("a" + "a") + ("a" + "hello")) + "") + "a") + "") + "")) + ("hello" + "")) + "\n\t"))) + "") + "foo bar") {
                case "a":
                    if (!(1 < 193)) {
                        var v23 int = -415
                        _ = v23
                    } else {
                        var v24 bool = (false || (-32768 == 0))
                        v25 := GenS1{F0: -32768}
                        if v24 {
                            v26 := v2[749]
                            var v27 uint32 = 2147483647
                            _ = v26
                            _ = v27
                        }
                        _ = v24
                        _ = v25
                    }
                case "\n\t":
                    v3 = v3
                }
                _ = v20
            }
        } else {
            var v28 uint64 = 18446744073709551615
            _ = v28
        }
    }
    if false {
        var v29 int8 = -97
        if true {
            for fi30 := 0; fi30 < 14; fi30++ {
                v4 = append(v4, (fi30 - 2147483647))
                _ = fi30
            }
        } else {
            var v31 bool = false
            switch 32768 {
            case -2147483648:
                v4 = append(v4, 32767)
            case 473:
                _ = len(v4)
            case 32767:
                v31 = true
            }
            v32 := v2[2]
            _ = v31
            _ = v32
        }
        _ = v29
    } else {
        for fi33 := 0; fi33 < 5; fi33++ {
            if false {
                v34 := GenS2{F0: 2, F1: false}
                _ = v34
            }
            var v35 uint16 = 255
            _ = fi33
            _ = v35
        }
        fmt.Println(v3)
        fmt.Println(v3)
    }
    var v36 string = ("a" + "a")
    v37 := GenS1{F0: -700}
    switch 642 {
    case 127:
        fmt.Println(v36, v36)
    case 42:
        v38 := GenS2{F0: 105, F1: true}
        _ = v38
    default:
        v39 := GenS2{F0: 2, F1: (((1 < -9223372036854775808) || false) && false)}
        _ = v39
    }
    v40 := GenS3{F0: -132}
    var v41 int8 = 0
    switch 0 {
    case 32768:
        var v42 int32 = -837
        _ = v42
    default:
        delete(v2, 4294967295)
    }
    switch -422 {
    case 9223372036854775807:
        fmt.Println(v36)
    }
    if (!true) {
        var v43 uint = 176
        _ = v43
    }
    v41 -= (^(v41))
    _ = len(v1)
    var v44 uint32 = 42
    _ = len(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v40
    _ = v41
    _ = v44
    _ = v5
}
