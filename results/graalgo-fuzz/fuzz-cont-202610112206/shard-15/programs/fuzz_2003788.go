package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool; F1 int8 }

func main() {
    var v1 bool = false
    v2 := []string{}
    v3 := make(map[uint32]string)
    v4 := []int{246, -734, -1, -9223372036854775808}
    var v5 uint16 = 65535
    v6 := []uint32{42}
    if v1 {
        v7 := v3[0]
        v8 := v3[4294967295]
        v3[619] = ((v8 + "hello") + v8)
        _ = v7
        _ = v8
    } else {
        var v9 uint16 = (v5 * (v5 + 931))
        if v1 {
            v3[1] = "\n\t"
            if (18446744073709551615 >= 2) {
                v6 = append(v6, 677)
                v10 := v3[2147483647]
                _ = v10
            }
        } else {
            for fi11 := 0; fi11 < 11; fi11++ {
                switch ("" + "") {
                case "a":
                    for fi12 := 0; fi12 < 7; fi12++ {
                        if ((42 > 1) && v1) {
                            var v13 uint = 2
                            _ = len(v4)
                            _ = v13
                        } else {
                            if false {
                                var v14 int16 = 42
                                v14 = (v14 ^ v14)
                                _ = v14
                            }
                            if (255 <= 0) {
                                if (1 <= -65) {
                                    if v1 {
                                        v1 = (!(322 != 2))
                                    } else {
                                        v15 := GenS1{F0: ("a" + ("hello" + "a")), F1: 0}
                                        _ = v15
                                    }
                                    v5 += v5
                                    v16 := v3[2]
                                    _ = v16
                                }
                            }
                            _ = len(v6)
                            v5 += v5
                            var v17 int = fi11
                            _ = v17
                        }
                        v18 := GenS2{F0: -32768}
                        v19 := GenS2{F0: 1}
                        _ = fi12
                        _ = v18
                        _ = v19
                    }
                }
                var v20 string = ""
                _ = fi11
                _ = v20
            }
            var v21 bool = v1
            _ = v21
        }
        _ = v9
    }
    if v1 {
        v5 -= 2
        v22 := v3[567]
        var v23 string = v22
        _ = v22
        _ = v23
    } else {
        var v24 int16 = -164
        for fi25 := 0; fi25 < 16; fi25++ {
            var v26 float64 = ((0.1 + 0.1) / 1.0)
            for fi27 := 0; fi27 < 7; fi27++ {
                v28 := GenS2{F0: v24}
                _ = fi27
                _ = v28
            }
            _ = fi25
            _ = v26
        }
        _ = v24
    }
    v29 := GenS3{F0: 4294967295}
    v6 = append(v6, uint32(v5))
    var v30 int64 = 0
    v3[61] = "\n\t"
    _ = len(v3)
    var v31 uint16 = 558
    v5 -= (v31 % ((32767) | 1))
    var v32 bool = v1
    v30 += (v30 << ((18446744073709551615) & 63))
    if ((v31 << ((255) & 63)) > (v5 | 1)) {
        switch (((((("" + "foo bar") + "a") + ("" + ("\n\t" + ((("a" + (("" + "a") + "")) + "hello") + "\n\t")))) + ("hello" + "a")) + "a") + (("foo bar" + "") + "a")) {
        case "foo bar":
            var v33 int32 = -2147483648
            _ = v33
        case "a":
            v31 -= 163
        }
        for fi34 := 0; fi34 < 19; fi34++ {
            fmt.Println(v1, v1, fi34)
            v6 = append(v6, 2147483647)
            var v35 uint32 = 2
            _ = fi34
            _ = v35
        }
    }
    for fi36 := 0; fi36 < 17; fi36++ {
        v2 = append(v2, "\n\t")
        v30 = (v30 >> ((464) & 63))
        if (255 < 255) {
            var v37 int32 = -1
            v32 = (4294967295 == 909)
            _ = v37
        } else {
            v38 := GenS4{F0: (!("" < "")), F1: 42}
            v32 = v1
            fmt.Println(v31, v31)
            _ = v38
        }
        _ = fi36
    }
    for fi39 := 0; fi39 < 20; fi39++ {
        v31 -= v31
        fmt.Println(v5, v1, v1)
        _ = fi39
    }
    if len(v6) > 0 {
        v6[0] = 409
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
