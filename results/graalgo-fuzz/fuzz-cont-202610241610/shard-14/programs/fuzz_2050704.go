package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 int64 }
type GenS2 struct { F0 uint; F1 uint64; F2 bool }

func main() {
    v1 := make(map[int]uint32)
    var v2 string = (("\n\t" + ("" + (("a" + ("a" + "")) + "\n\t"))) + "a")
    var v3 int64 = -1
    var v4 string = (v2 + (v2 + v2))
    if ((!true) || false) {
        for fi5 := 0; fi5 < 19; fi5++ {
            var v6 uint16 = 467
            var v7 int64 = -9223372036854775808
            _ = fi5
            _ = v6
            _ = v7
        }
        fmt.Println(v2, v4)
        switch 4294967295 {
        case 4294967295:
            var v8 uint64 = 18446744073709551615
            _ = v8
        case -727:
            switch v4 {
            case "hello":
                if (!(0.1 == 1e10)) {
                    var v9 uint64 = 18446744073709551615
                    for fi10 := 0; fi10 < 11; fi10++ {
                        v11 := GenS1{F0: 0, F1: (42 == 100), F2: 0}
                        _ = fi10
                        _ = v11
                    }
                    _ = v9
                } else {
                    v2 = v4
                    if true {
                        for fi12 := 0; fi12 < 11; fi12++ {
                            v2 = "hello"
                            if false {
                                if (32767 != fi12) {
                                    var v13 uint = 826
                                    for fi14 := 0; fi14 < 2; fi14++ {
                                        fmt.Println(fi12)
                                        var v15 int64 = int64(v13)
                                        switch (^(fi12)) {
                                        case 2147483647:
                                            _ = len(v1)
                                        case -9223372036854775808:
                                            v16 := []int{}
                                            _ = v16
                                        }
                                        v15 = (v15 * (v3 % (((v3 ^ (v15 * (v15 / ((79) | 1))))) | 1)))
                                        _ = fi14
                                        _ = v15
                                    }
                                    var v17 uint = v13
                                    _ = v13
                                    _ = v17
                                } else {
                                    if (fi12 > (fi12 & 319)) {
                                        if (868 >= 0) {
                                            var v18 int32 = -1
                                            var v19 string = "a"
                                            for fi20 := 0; fi20 < 2; fi20++ {
                                                fmt.Println(fi20, v4)
                                                delete(v1, (fi12 & fi12))
                                                _ = fi20
                                            }
                                            _ = v18
                                            _ = v19
                                        }
                                        v21 := v1[(fi12 * fi12)]
                                        _ = v21
                                    }
                                    if true {
                                        for fi22 := 0; fi22 < 18; fi22++ {
                                            var v23 bool = ((128 > 1) && true)
                                            v24 := []int{}
                                            v1[(fi12 ^ (fi12 >> ((65535) & 63)))] = 236
                                            _ = fi22
                                            _ = v23
                                            _ = v24
                                        }
                                        if true {
                                            var v25 string = ""
                                            v25 = ((v2 + v2) + v4)
                                            fmt.Println(v25)
                                            _ = v25
                                        }
                                        var v26 int16 = 42
                                        _ = v26
                                    } else {
                                        var v27 int16 = 0
                                        _ = v27
                                    }
                                }
                                var v28 uint64 = 2
                                _ = v28
                            } else {
                                v2 = (v4 + ((("\n\t" + (v4 + v4)) + v2) + (v4 + v2)))
                                var v29 uint64 = 2
                                v30 := v1[128]
                                _ = v29
                                _ = v30
                            }
                            _ = fi12
                        }
                        v31 := v1[933]
                        v32 := GenS2{F0: 1, F1: 2, F2: false}
                        _ = v31
                        _ = v32
                    }
                    v33 := GenS1{F0: 255, F1: (false || true), F2: v3}
                    _ = v33
                }
            case "foo bar":
                fmt.Println(v4, v2)
            }
        }
    }
    fmt.Println(v2, v2, v2)
    _ = len(v1)
    for fi34 := 0; fi34 < 14; fi34++ {
        fmt.Println(fi34, v3)
        var v35 uint32 = 2
        _ = fi34
        _ = v35
    }
    v3 += v3
    v3 = v3
    delete(v1, 2147483647)
    if true {
        if (42 < 42) {
            v36 := GenS1{F0: 297, F1: (true && (0.0 != 1e10)), F2: (^(v3))}
            _ = v36
        } else {
            var v37 uint8 = 0
            v2 = v2
            if (true && (!true)) {
                if true {
                    var v38 uint64 = 2
                    _ = v38
                }
            }
            _ = v37
        }
    } else {
        var v39 uint16 = 1
        for fi40 := 0; fi40 < 7; fi40++ {
            for fi41 := 0; fi41 < 6; fi41++ {
                var v42 int32 = 1
                v3 = v3
                _ = fi41
                _ = v42
            }
            _ = fi40
        }
        var v43 string = v2
        _ = v39
        _ = v43
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
