package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 int64; F2 bool }

func main() {
    v1 := []int{-238, 254, -286}
    var v2 uint64 = 18446744073709551615
    var v3 uint = 18446744073709551615
    v4 := make(map[string]string)
    switch 787 {
    case 4294967295:
        var v5 uint32 = 0
        _ = v5
    case 9223372036854775807:
        if false {
            v6 := v4["\n\t"]
            _ = v6
        } else {
            fmt.Println(v3, v3)
            v7 := GenS1{F0: 32768}
            _ = v7
        }
    }
    switch 42 {
    case 42:
        for fi8 := 0; fi8 < 4; fi8++ {
            fmt.Println(v3, v3, v3)
            var v9 uint64 = (v2 << ((v3) & 63))
            for fi10 := 0; fi10 < 5; fi10++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v9 += (v2 >> ((544) & 63))
                var v11 int16 = -36
                _ = fi10
                _ = v11
            }
            _ = fi8
            _ = v9
        }
    case 32768:
        for fi12 := 0; fi12 < 16; fi12++ {
            if (!(2147483647 <= 2147483647)) {
                fmt.Println(v2, v2, v2)
                if true {
                    for fi13 := 0; fi13 < 12; fi13++ {
                        v2 += (v2 >> (((v3 >> ((0) & 63))) & 63))
                        _ = fi13
                    }
                }
            }
            switch "foo bar" {
            case "":
                v3 -= v3
            }
            _ = fi12
        }
    }
    v3 = (v3 - v3)
    v14 := GenS1{F0: 0}
    v15 := GenS1{F0: -668}
    var v16 uint8 = 42
    fmt.Println(v3)
    if (!(0 >= -88)) {
        if true {
            _ = len(v4)
            if true {
                var v17 int = -829
                v18 := GenS1{F0: (v17 - v17)}
                switch (-(v17)) {
                case 32768:
                    v19 := GenS1{F0: v17}
                    _ = v19
                case 2147483647:
                    if true {
                        fmt.Println(v17, v3, v2)
                        v20 := GenS1{F0: v17}
                        for fi21 := 0; fi21 < 6; fi21++ {
                            delete(v4, "hello")
                            var v22 uint16 = 32767
                            _ = fi21
                            _ = v22
                        }
                        _ = v20
                    } else {
                        var v23 int = (v17 + v17)
                        v24 := GenS2{F0: 42, F1: 5, F2: (62 < -64)}
                        _ = v23
                        _ = v24
                    }
                case 323:
                    for fi25 := 0; fi25 < 3; fi25++ {
                        switch (v17 >> (((v3 % ((v3) | 1))) & 63)) {
                        case -2147483648:
                            var v26 uint32 = 153
                            _ = v26
                        }
                        var v27 string = "foo bar"
                        _ = fi25
                        _ = v27
                    }
                default:
                    if (42 <= -114) {
                        v17 += v17
                    } else {
                        v17 = (v17 - v17)
                    }
                }
                _ = v17
                _ = v18
            } else {
                switch "foo bar" {
                case "":
                    switch ("hello" + "foo bar") {
                    case "":
                        if true {
                            if len(v1) > 0 {
                                v1[0] = 127
                            }
                        } else {
                            for fi28 := 0; fi28 < 16; fi28++ {
                                for fi29 := 0; fi29 < 17; fi29++ {
                                    fmt.Println(v16, fi28, v16)
                                    _ = fi29
                                }
                                _ = fi28
                            }
                            v16 += 2
                        }
                    case "\n\t":
                        v3 += v3
                    }
                case "hello":
                    for fi30 := 0; fi30 < 12; fi30++ {
                        fmt.Println(fi30)
                        _ = fi30
                    }
                default:
                    v16 = v16
                }
                v2 = (^(v2))
            }
            var v31 float32 = (1e10 - 0.0)
            _ = v31
        } else {
            var v32 uint16 = uint16(v16)
            var v33 bool = false
            var v34 bool = true
            _ = v32
            _ = v33
            _ = v34
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
