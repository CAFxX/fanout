package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8; F1 int16; F2 int64 }
func genFn0(p0 string) uint64 {
    switch 127 {
    case 128:
        p0 = "\n\t"
    }
    var v6 int64 = -9223372036854775808
    _ = p0
    _ = v6
    return 2
}

func main() {
    var v1 uint8 = 42
    v2 := make(map[string]uint32)
    var v3 int32 = 2147483647
    v4 := make(map[int]int)
    v5 := make(map[int]int)
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 string = "foo bar"
        _ = fi7
        _ = v8
    }
    v1 -= 1
    if false {
        var v9 uint32 = 2147483647
        _ = v9
    }
    fmt.Println(v3, v1)
    v3 += (-(v3))
    var v10 uint8 = 137
    var v11 int32 = (v3 >> ((738) & 63))
    for fi12 := 0; fi12 < 15; fi12++ {
        v13 := GenS1{F0: -1}
        if true {
            for fi14 := 0; fi14 < 4; fi14++ {
                delete(v5, 127)
                fmt.Println(fi12)
                _ = fi14
            }
            fmt.Println(v3)
        }
        var v15 int32 = (v3 % (((^(v3))) | 1))
        _ = fi12
        _ = v13
        _ = v15
    }
    v16 := []int{1}
    var v17 float64 = ((float64(0.0) / 1.0) + 3.14159)
    for fi18 := 0; fi18 < 9; fi18++ {
        if (4294967295 == 2147483648) {
            v3 = 1
            if (true && false) {
                var v19 int16 = -32768
                for fi20 := 0; fi20 < 17; fi20++ {
                    var v21 float32 = 0.0
                    fmt.Println(v3)
                    v22 := v4[int(v11)]
                    var v23 float32 = (v21 - 3.14159)
                    _ = fi20
                    _ = v21
                    _ = v22
                    _ = v23
                }
                v24 := GenS2{F0: 9223372036854775807}
                _ = v19
                _ = v24
            }
            for fi25 := 0; fi25 < 7; fi25++ {
                v26 := v5[(fi18 / ((9223372036854775807) | 1))]
                _ = len(v4)
                _ = fi25
                _ = v26
            }
        } else {
            if true {
                var v27 int32 = (^(v3))
                if false {
                    switch (-(fi18)) {
                    case 127:
                        for fi28 := 0; fi28 < 10; fi28++ {
                            var v29 uint16 = 32767
                            if len(v16) > 0 {
                                v16[0] = fi28
                            }
                            v3 = v11
                            v10 = v10
                            _ = fi28
                            _ = v29
                        }
                    case 2147483647:
                        var v30 int16 = 512
                        _ = v30
                    case 367:
                        switch (("a" + "hello") + "a") {
                        case "foo bar":
                            v27 = v11
                        }
                    }
                }
                v31 := GenS1{F0: 1}
                _ = v27
                _ = v31
            }
            if false {
                var v32 int16 = 1
                v3 += -2147483648
                _ = v32
            } else {
                fmt.Println(v1)
                v33 := GenS2{F0: 42}
                _ = v33
            }
        }
        for fi34 := 0; fi34 < 18; fi34++ {
            var v35 uint = 4294967295
            if false {
                var v36 int64 = 1
                v35 = v35
                _ = v36
            }
            var v37 uint16 = 32767
            _ = fi34
            _ = v35
            _ = v37
        }
        _ = fi18
    }
    if true {
        switch "\n\t" {
        case "a":
            delete(v5, 684)
        case "foo bar":
            v3 += 2147483647
        case "hello":
            if (1 <= 42) {
                if (v17 < (v17 + (1e10 + 1e10))) {
                    switch "hello" {
                    case "a":
                        fmt.Println(v3, v1)
                    case "hello":
                        var v38 int8 = -23
                        _ = v38
                    default:
                        v39 := v2["\n\t"]
                        _ = v39
                    }
                    v40 := GenS1{F0: 9223372036854775807}
                    _ = v40
                } else {
                    v17 += (1e10 + v17)
                    v10 = (v10 + 36)
                }
            } else {
                if true {
                    var v41 int16 = 88
                    var v42 uint = 2
                    _ = v41
                    _ = v42
                } else {
                    v10 = (v10 | 42)
                    var v43 uint = 65535
                    var v44 uint16 = 32767
                    _ = v43
                    _ = v44
                }
                v10 = (v10 * v1)
            }
        default:
            for fi45 := 0; fi45 < 13; fi45++ {
                var v46 int64 = 0
                v47 := GenS1{F0: 1}
                _ = fi45
                _ = v46
                _ = v47
            }
        }
    }
    v48 := GenS3{F0: (^(v10)), F1: 1000, F2: int64(v11)}
    v49 := GenS2{F0: 9223372036854775807}
    delete(v4, 280)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
}
