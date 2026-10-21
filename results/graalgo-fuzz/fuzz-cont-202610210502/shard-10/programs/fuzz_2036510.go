package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32 }
func genFn0(p0 int8, p1 int16) string {
    if (20 <= 1) {
        var v7 uint32 = 1
        _ = v7
    } else {
        var v8 uint64 = 2
        _ = v8
    }
    var v9 int = -2147483648
    for fi10 := 0; fi10 < 2; fi10++ {
        fmt.Println(p0)
        for fi11 := 0; fi11 < 6; fi11++ {
            var v12 int16 = 132
            _ = fi11
            _ = v12
        }
        fmt.Println(p0)
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = v9
    return (("hello" + ("foo bar" + ("" + "hello"))) + "hello")
}
func genFn1(p0 int64, p1 uint32) int32 {
    p0 = (p0 & (p0 % (((p0 % ((p0) | 1))) | 1)))
    v13 := []int{}
    var v14 int32 = 1
    for fi15 := 0; fi15 < 19; fi15++ {
        fmt.Println(fi15, v14)
        v13 = append(v13, fi15)
        _ = fi15
    }
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return 883
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]uint32)
    v3 := make(map[int]int)
    var v4 bool = true
    var v5 int = 256
    var v6 int32 = 0
    v5 = (v5 | (v5 % (((^(v5))) | 1)))
    v5 = 127
    _ = len(v2)
    if true {
        var v16 int64 = 0
        _ = v16
    } else {
        var v17 uint16 = 1
        _ = v17
    }
    if v4 {
        var v18 string = "a"
        for fi19 := 0; fi19 < 13; fi19++ {
            switch (v5 >> ((18446744073709551615) & 63)) {
            case 950:
                for fi20 := 0; fi20 < 4; fi20++ {
                    var v21 uint64 = 2
                    _ = fi20
                    _ = v21
                }
            case 127:
                for fi22 := 0; fi22 < 12; fi22++ {
                    var v23 uint16 = 2
                    var v24 uint64 = 42
                    fmt.Println(v18)
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            default:
                var v25 uint32 = 0
                _ = v25
            }
            _ = fi19
        }
        v4 = (2 == 190)
        _ = v18
    }
    v26 := GenS1{F0: "\n\t"}
    v27 := []int{(v5 >> ((42) & 63)), v5, (v5 + v5)}
    v28 := GenS2{F0: v6}
    if v4 {
        switch ((((("" + "\n\t") + "foo bar") + "foo bar") + ("hello" + "a")) + "\n\t") {
        case "hello":
            switch 128 {
            case -2147483648:
                v29 := v1["foo bar"]
                _ = v29
            case 32768:
                for fi30 := 0; fi30 < 2; fi30++ {
                    v4 = v4
                    v31 := GenS2{F0: v6}
                    _ = fi30
                    _ = v31
                }
            }
        }
    }
    var v32 string = "foo bar"
    v33 := GenS1{F0: (v32 + "a")}
    for fi34 := 0; fi34 < 7; fi34++ {
        v32 = "\n\t"
        _ = len(v3)
        _ = fi34
    }
    if (2 >= 436) {
        switch ("foo bar" + v32) {
        case "\n\t":
            for fi35 := 0; fi35 < 14; fi35++ {
                switch ((v32 + "\n\t") + "\n\t") {
                case "foo bar":
                    if (!v4) {
                        if (0 > -931) {
                            fmt.Println(fi35, v6, v4)
                        } else {
                            fmt.Println(v32)
                            v4 = ((float64((0.0 + (1e10 * (0.1 + 1.5)))) / -2.25) < 0.1)
                        }
                        var v36 uint8 = 53
                        _ = v36
                    }
                default:
                    switch (v5 % ((fi35) | 1)) {
                    case 127:
                        v27 = append(v27, 127)
                    case 1:
                        var v37 uint = 1
                        _ = v37
                    case -9223372036854775808:
                        for fi38 := 0; fi38 < 15; fi38++ {
                            fmt.Println(v5, fi35)
                            _ = fi38
                        }
                    }
                }
                v5 -= (-(fi35))
                _ = fi35
            }
        case "foo bar":
            v32 = v32
        case "":
            for fi39 := 0; fi39 < 13; fi39++ {
                var v40 int32 = -2147483648
                _ = fi39
                _ = v40
            }
        }
    } else {
        var v41 uint32 = 159
        switch v5 {
        case -959:
            switch v5 {
            case -671:
                v41 -= 65535
            default:
                switch (v5 % ((v5) | 1)) {
                case 4294967295:
                    v6 = (v6 / ((v6) | 1))
                default:
                    for fi42 := 0; fi42 < 18; fi42++ {
                        fmt.Println(v32, v5, v5)
                        delete(v1, "foo bar")
                        if (false || (2 >= 0)) {
                            v5 = v5
                            _ = len(v1)
                        }
                        _ = fi42
                    }
                }
            }
        case -924:
            var v43 uint32 = (v41 - 2147483647)
            _ = v43
        }
        _ = v41
    }
    v5 = (v5 >> ((2) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
