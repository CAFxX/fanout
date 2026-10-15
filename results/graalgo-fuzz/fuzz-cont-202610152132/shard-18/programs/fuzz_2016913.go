package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 int8 }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 bool; F1 uint32 }
type GenS5 struct { F0 int16; F1 uint16 }
type GenS6 struct { F0 int64; F1 bool; F2 uint }
func genFn0(p0 int64) string {
    v7 := []int{2147483647, 1, 2147483647}
    var v8 bool = true
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    for fi9 := 0; fi9 < 16; fi9++ {
        fmt.Println(fi9)
        v10 := GenS1{F0: 80, F1: ("a" + "a"), F2: -119}
        _ = fi9
        _ = v10
    }
    if v8 {
        var v11 uint8 = 0
        _ = len(v7)
        _ = v11
    }
    _ = p0
    _ = v7
    _ = v8
    return "foo bar"
}

func main() {
    v1 := make(map[string]bool)
    var v2 int16 = 1
    var v3 int8 = -20
    var v4 bool = (-9223372036854775808 > -9223372036854775808)
    v5 := make(map[string]bool)
    v6 := make(map[uint32]bool)
    for fi12 := 0; fi12 < 16; fi12++ {
        switch "" {
        case "a":
            v13 := []int{fi12, (fi12 ^ (fi12 + (fi12 + (fi12 << ((1) & 63)))))}
            _ = v13
        default:
            v14 := []int{(-(fi12)), fi12}
            _ = v14
        }
        for fi15 := 0; fi15 < 8; fi15++ {
            if (!(!v4)) {
                v16 := GenS1{F0: 2, F1: "hello", F2: 100}
                v17 := GenS2{F0: v2, F1: v2}
                v18 := v1["\n\t"]
                _ = v16
                _ = v17
                _ = v18
            } else {
                v19 := GenS1{F0: 0, F1: "foo bar", F2: v3}
                _ = v19
            }
            var v20 string = "a"
            _ = fi15
            _ = v20
        }
        v5["foo bar"] = v4
        _ = fi12
    }
    v2 += -1
    if v4 {
        delete(v6, 910)
        v3 = v3
    } else {
        if v4 {
            fmt.Println(v3, v2)
        }
        for fi21 := 0; fi21 < 18; fi21++ {
            var v22 int64 = 595
            v23 := []int{(^(fi21)), (fi21 ^ fi21), fi21, fi21}
            _ = fi21
            _ = v22
            _ = v23
        }
        if (!((-2147483648 > 340) || (9223372036854775807 == 42))) {
            switch (("a" + "a") + ("foo bar" + "\n\t")) {
            case "foo bar":
                for fi24 := 0; fi24 < 1; fi24++ {
                    var v25 uint64 = 321
                    var v26 int = fi24
                    switch (("a" + "foo bar") + "foo bar") {
                    case "foo bar":
                        if (42 <= 2147483647) {
                            v4 = false
                            var v27 uint16 = 0
                            _ = v27
                        } else {
                            v2 += v2
                            fmt.Println(v25)
                        }
                    }
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            case "":
                var v28 uint32 = 4294967295
                _ = v28
            case "a":
                var v29 string = ""
                _ = v29
            default:
                v2 = v2
            }
        }
    }
    v6[4294967295] = ((v2 ^ (v2 * -1)) <= (v2 * 42))
    switch 2147483647 {
    case 886:
        if (-273 >= 9223372036854775807) {
            v4 = (v3 <= 127)
            switch -1 {
            case 42:
                if v4 {
                    var v30 int16 = (v2 / ((-32768) | 1))
                    v31 := GenS1{F0: 128, F1: "hello", F2: v3}
                    _ = v30
                    _ = v31
                }
            case 127:
                switch ("hello" + "") {
                case "foo bar":
                    switch 32767 {
                    case 0:
                        var v32 float32 = ((-2.25 * (0.0 / float32(-2.25))) - ((0.0 + (3.14159 + 1.5)) + 0.0))
                        _ = v32
                    }
                }
            case -565:
                v33 := GenS3{F0: 187}
                _ = v33
            default:
                v2 += 32767
            }
            for fi34 := 0; fi34 < 8; fi34++ {
                v4 = v4
                _ = fi34
            }
        }
    case 368:
        delete(v1, "a")
    case 4294967295:
        v4 = (0 != 68)
    default:
        var v35 uint8 = 127
        _ = v35
    }
    v36 := []int{}
    v3 = (-(v3))
    for fi37 := 0; fi37 < 3; fi37++ {
        var v38 int64 = 0
        for fi39 := 0; fi39 < 20; fi39++ {
            v38 = v38
            if (v4 || (32767 >= 0)) {
                v40 := GenS4{F0: (2 >= 2), F1: 561}
                _ = v40
            } else {
                if (-514 != -2147483648) {
                    switch (fi39 + fi37) {
                    case -698:
                        var v41 int = 9223372036854775807
                        _ = v41
                    case 1:
                        fmt.Println(v2, v2, v2)
                    }
                } else {
                    _ = len(v36)
                    v2 += v2
                    fmt.Println(v4, v4)
                    fmt.Println(fi39, v2, fi39)
                    _ = len(v36)
                    var v42 bool = (true && true)
                    v43 := GenS5{F0: (v2 << ((42) & 63)), F1: 255}
                    _ = v42
                    _ = v43
                }
            }
            _ = fi39
        }
        _ = fi37
        _ = v38
    }
    var v44 uint64 = 18446744073709551615
    v45 := GenS6{F0: -452, F1: v4, F2: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
