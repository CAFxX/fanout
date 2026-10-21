package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 uint32; F2 uint16 }
func genFn0(p0 int16) int {
    p0 += (p0 / ((42) | 1))
    var v6 int16 = p0
    _ = p0
    _ = v6
    return 662
}

func main() {
    var v1 bool = (32767 > 42)
    v2 := make(map[uint32]int)
    var v3 float64 = 1.5
    var v4 string = ("foo bar" + "a")
    var v5 uint32 = 2147483647
    v5 = (v5 + (v5 % (((v5 % ((v5) | 1))) | 1)))
    v5 -= (v5 & 2147483647)
    v5 = v5
    delete(v2, v5)
    var v7 int8 = 42
    v1 = v1
    fmt.Println(v4, v7)
    v8 := v2[2147483647]
    if false {
        if false {
            var v9 bool = v1
            _ = v9
        } else {
            for fi10 := 0; fi10 < 11; fi10++ {
                v11 := v2[v5]
                for fi12 := 0; fi12 < 17; fi12++ {
                    var v13 int8 = v7
                    var v14 int = -1
                    v15 := v2[(v5 - (v5 - 997))]
                    fmt.Println(v1, v4)
                    if v1 {
                        v1 = (!v1)
                        if (!v1) {
                            v16 := GenS1{F0: 0, F1: 0}
                            var v17 uint16 = 609
                            fmt.Println(v7)
                            v18 := GenS2{F0: 0}
                            switch v4 {
                            case "hello":
                                if (v1 || v1) {
                                    if v1 {
                                        if v1 {
                                            fmt.Println(v4, v8, v14)
                                            var v19 float32 = 1.5
                                            v5 += (^(v5))
                                            _ = v19
                                        }
                                        fmt.Println(v11)
                                        v14 = int(v7)
                                    } else {
                                        switch fi12 {
                                        case -967:
                                            v4 = v4
                                        }
                                        v1 = v1
                                        v13 -= v7
                                    }
                                    v20 := []int{(v8 + (v8 - genFn0(-934))), fi12, (v14 & fi10)}
                                    _ = v20
                                }
                            case "a":
                                switch (fi12 << ((65535) & 63)) {
                                case -934:
                                    var v21 int16 = 1000
                                    _ = v21
                                }
                            case "foo bar":
                                v22 := GenS2{F0: 0}
                                _ = v22
                            }
                            var v23 float64 = v3
                            v15 = 2147483647
                            _ = v16
                            _ = v17
                            _ = v18
                            _ = v23
                        }
                    } else {
                        fmt.Println(v3, v11)
                    }
                    _ = fi12
                    _ = v13
                    _ = v14
                    _ = v15
                }
                _ = fi10
                _ = v11
            }
        }
        v24 := []int{(v8 ^ genFn0(163))}
        _ = v24
    } else {
        var v25 uint32 = 290
        if v1 {
            var v26 int = v8
            _ = len(v2)
            _ = v26
        } else {
            for fi27 := 0; fi27 < 19; fi27++ {
                v2[(v5 >> ((55) & 63))] = 830
                for fi28 := 0; fi28 < 13; fi28++ {
                    var v29 uint32 = 262
                    fmt.Println(v29)
                    var v30 int16 = 514
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                _ = fi27
            }
            var v31 int32 = 2147483647
            if (32767 == 1) {
                var v32 int16 = -32768
                for fi33 := 0; fi33 < 7; fi33++ {
                    v34 := GenS3{F0: genFn0(v32), F1: (^(v5)), F2: 886}
                    switch 128 {
                    case 4294967295:
                        v35 := GenS2{F0: 2}
                        _ = v35
                    case 127:
                        fmt.Println(v7)
                    case 42:
                        v4 = v4
                    }
                    var v36 int32 = (^(v31))
                    _ = fi33
                    _ = v34
                    _ = v36
                }
                _ = v32
            }
            _ = v31
        }
        var v37 uint = 706
        _ = v25
        _ = v37
    }
    v8 += v8
    switch "\n\t" {
    case "\n\t":
        var v38 int32 = 2147483647
        _ = v38
    }
    var v39 uint8 = 1
    v39 += v39
    for fi40 := 0; fi40 < 20; fi40++ {
        if false {
            if ((!(9223372036854775807 == int64(fi40))) || ((v8 & (-(v8))) <= genFn0(42))) {
                var v41 bool = false
                if false {
                    v42 := []int{42, 9223372036854775807, (-(v8)), 229}
                    _ = v42
                }
                var v43 uint16 = 65535
                _ = v41
                _ = v43
            }
            if (9223372036854775807 < 42) {
                fmt.Println(v3, v5, fi40)
                if v1 {
                    v5 += v5
                }
            }
        } else {
            var v44 int64 = -772
            var v45 float64 = 0.0
            _ = v44
            _ = v45
        }
        _ = fi40
    }
    v46 := []int{genFn0(0), (v8 - genFn0(32767)), 1}
    v47 := GenS3{F0: 538, F1: v5, F2: 32767}
    v39 -= 26
    var v48 string = v4
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v7
    _ = v8
}
