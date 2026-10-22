package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16; F1 uint }
func genFn0(p0 int) int16 {
    switch "foo bar" {
    case "a":
        var v7 uint8 = 255
        _ = v7
    case "foo bar":
        switch ("" + "foo bar") {
        case "a":
            var v8 int16 = -32768
            _ = v8
        }
    case "hello":
        p0 += (p0 | p0)
    }
    var v9 int8 = -61
    _ = p0
    _ = v9
    return -32768
}

func main() {
    v1 := make(map[int]int)
    var v2 uint = 65535
    v3 := make(map[string]bool)
    v4 := make(map[string]string)
    var v5 uint = v2
    var v6 bool = true
    v10 := GenS1{F0: -684}
    switch "foo bar" {
    case "a":
        if v6 {
            v11 := GenS1{F0: -1}
            if v6 {
                var v12 int32 = -2147483648
                if (-9223372036854775808 >= 1) {
                    v6 = (0 != -517)
                    for fi13 := 0; fi13 < 1; fi13++ {
                        var v14 int16 = genFn0(fi13)
                        if true {
                            for fi15 := 0; fi15 < 16; fi15++ {
                                if (232 > 93) {
                                    v14 += v14
                                    if (!false) {
                                        switch "foo bar" {
                                        case "hello":
                                            v6 = (-37 > 43)
                                        }
                                    }
                                } else {
                                    var v16 uint64 = uint64(v5)
                                    _ = v16
                                }
                                var v17 int8 = -36
                                var v18 uint32 = 0
                                _ = fi15
                                _ = v17
                                _ = v18
                            }
                            v19 := []int{42, (^(fi13)), (fi13 ^ fi13), fi13}
                            _ = v19
                        }
                        _ = fi13
                        _ = v14
                    }
                    v20 := GenS2{F0: 101, F1: 0}
                    _ = v20
                } else {
                    var v21 uint16 = 255
                    if (0.0 <= ((1.5 + 1e10) * (3.14159 * 1e10))) {
                        v22 := v4["hello"]
                        var v23 float32 = 3.14159
                        _ = v22
                        _ = v23
                    }
                    if v6 {
                        v5 = (v2 | 18446744073709551615)
                        v6 = ("hello" > "\n\t")
                    }
                    _ = v21
                }
                _ = v12
            } else {
                for fi24 := 0; fi24 < 8; fi24++ {
                    for fi25 := 0; fi25 < 14; fi25++ {
                        switch "" {
                        case "":
                            v26 := GenS2{F0: 240, F1: 42}
                            _ = v26
                        }
                        _ = fi25
                    }
                    _ = fi24
                }
            }
            var v27 int32 = 1
            _ = v11
            _ = v27
        } else {
            var v28 int32 = 183
            var v29 bool = (!(true && (0 >= v28)))
            _ = v28
            _ = v29
        }
    case "foo bar":
        v2 += (v2 % ((v2) | 1))
    default:
        v2 -= (^(v5))
    }
    v6 = v6
    v30 := GenS1{F0: 0}
    v31 := GenS1{F0: -6}
    switch (("\n\t" + "foo bar") + ("" + "")) {
    case "hello":
        delete(v1, 167)
    }
    v32 := GenS3{F0: 1}
    v2 += 255
    for fi33 := 0; fi33 < 3; fi33++ {
        v5 -= 2
        for fi34 := 0; fi34 < 13; fi34++ {
            switch 9223372036854775807 {
            case 4294967295:
                var v35 int8 = 99
                _ = v35
            case 1:
                var v36 uint16 = 65535
                _ = v36
            }
            if v6 {
                v37 := GenS4{F0: 68, F1: v5}
                fmt.Println(v5)
                _ = v37
            }
            v2 -= (v5 ^ 255)
            _ = fi34
        }
        var v38 int16 = genFn0((fi33 % (((fi33 & fi33)) | 1)))
        _ = fi33
        _ = v38
    }
    for fi39 := 0; fi39 < 16; fi39++ {
        var v40 uint8 = 0
        _ = fi39
        _ = v40
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
