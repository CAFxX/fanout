package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int16, p1 int) bool {
    switch (p1 / ((p1) | 1)) {
    case -1:
        var v7 uint = 957
        _ = v7
    case -9223372036854775808:
        p0 += p0
    default:
        var v8 int64 = 0
        _ = v8
    }
    v9 := []int{42, 127, p1}
    for fi10 := 0; fi10 < 8; fi10++ {
        v9 = append(v9, (p1 + -2147483648))
        var v11 uint64 = 884
        var v12 uint = 2
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v9
    return ((42 != -1) || (2147483647 >= 0))
}
func genFn1() int16 {
    v13 := GenS1{F0: 1}
    v14 := []int{0, -916, 9223372036854775807}
    var v15 float32 = (1e10 - 3.14159)
    var v16 uint32 = 1
    v14 = append(v14, -698)
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    return 374
}

func main() {
    var v1 int = 32768
    var v2 int64 = 9223372036854775807
    var v3 uint = 632
    v4 := make(map[string]bool)
    var v5 int32 = 0
    var v6 uint16 = 0
    for fi17 := 0; fi17 < 8; fi17++ {
        v3 -= (v3 * 0)
        _ = fi17
    }
    v18 := v4["hello"]
    switch (v1 - v1) {
    case 127:
        v1 = v1
    case 32767:
        var v19 uint32 = 0
        _ = v19
    case 968:
        switch (v1 / ((0) | 1)) {
        case -43:
            v18 = v18
        case 2147483647:
            if true {
                var v20 int8 = int8(v1)
                switch v1 {
                case -9223372036854775808:
                    var v21 int16 = 864
                    _ = v21
                default:
                    v1 += -2147483648
                }
                v6 = v6
                _ = v20
            }
        case 42:
            for fi22 := 0; fi22 < 18; fi22++ {
                for fi23 := 0; fi23 < 20; fi23++ {
                    v1 += v1
                    _ = len(v4)
                    switch fi23 {
                    case 32767:
                        v24 := GenS2{F0: v6}
                        _ = v24
                    case 42:
                        v25 := GenS2{F0: v6}
                        _ = v25
                    default:
                        var v26 uint = v3
                        _ = v26
                    }
                    _ = fi23
                }
                _ = fi22
            }
        }
    default:
        if false {
            var v27 int8 = int8(v1)
            _ = v27
        } else {
            if v18 {
                var v28 int64 = 1
                fmt.Println(v18, v1, v28)
                _ = v28
            } else {
                switch 136 {
                case 2147483647:
                    var v29 int32 = v5
                    _ = v29
                }
                if v18 {
                    var v30 uint64 = 4294967295
                    _ = v30
                } else {
                    v2 += v2
                    if ((^(v5)) > (-(v5))) {
                        for fi31 := 0; fi31 < 4; fi31++ {
                            switch 32768 {
                            case 2147483647:
                                v18 = (208 <= 108)
                            }
                            _ = fi31
                        }
                        fmt.Println(v1, v1)
                        for fi32 := 0; fi32 < 7; fi32++ {
                            for fi33 := 0; fi33 < 12; fi33++ {
                                var v34 uint32 = 2147483647
                                switch (("\n\t" + ((("foo bar" + "foo bar") + "") + "")) + ("\n\t" + ("" + "\n\t"))) {
                                case "foo bar":
                                    v35 := GenS2{F0: (v6 + uint16(v34))}
                                    _ = v35
                                case "\n\t":
                                    switch 0 {
                                    case 32768:
                                        fmt.Println(fi32, v6)
                                        fmt.Println(v6)
                                    case -234:
                                        var v36 int64 = (v2 - int64(v34))
                                        _ = v36
                                    case -617:
                                        _ = len(v4)
                                    default:
                                        v37 := []int{v1, (fi32 % ((-9223372036854775808) | 1))}
                                        _ = v37
                                    }
                                default:
                                    var v38 uint32 = (v34 + v34)
                                    _ = v38
                                }
                                v39 := []int{}
                                v40 := GenS1{F0: 127}
                                var v41 int64 = v2
                                v1 -= fi32
                                v1 = (v1 << ((v3) & 63))
                                v34 = 4294967295
                                _ = fi33
                                _ = v34
                                _ = v39
                                _ = v40
                                _ = v41
                            }
                            var v42 int8 = 0
                            var v43 uint = v3
                            _ = fi32
                            _ = v42
                            _ = v43
                        }
                    }
                    v3 = (v3 % (((^(v3))) | 1))
                }
            }
        }
    }
    v18 = ((4294967295 >= 917) && v18)
    if ((float64((-2.25 * 0.0)) >= -2.25) || (!(1 != v5))) {
        v18 = v18
    }
    v6 = (v6 & v6)
    var v44 int16 = 726
    if (("\n\t" + ("" + ("a" + "\n\t"))) > "\n\t") {
        v3 += 18446744073709551615
    }
    var v45 float32 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
