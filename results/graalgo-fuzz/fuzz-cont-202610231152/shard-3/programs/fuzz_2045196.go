package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 uint32; F2 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint }
func genFn0() int16 {
    var v7 string = "\n\t"
    v8 := []int{}
    fmt.Println(v7)
    var v9 int64 = -875
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 bool = false
        if v11 {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            v8 = append(v8, -880)
            for fi12 := 0; fi12 < 3; fi12++ {
                var v13 uint32 = 597
                _ = fi12
                _ = v13
            }
        }
        _ = fi10
        _ = v11
    }
    _ = v7
    _ = v8
    _ = v9
    return -414
}
func genFn1(p0 uint64) int8 {
    var v14 int = int(p0)
    p0 -= (p0 * p0)
    p0 += (p0 ^ (^(p0)))
    v15 := GenS1{F0: 1}
    p0 = p0
    _ = p0
    _ = v14
    _ = v15
    return 100
}

func main() {
    var v1 int = 32767
    var v2 uint = 882
    var v3 int8 = -1
    v4 := []string{"foo bar", ("\n\t" + (("" + "a") + ("foo bar" + "hello")))}
    v5 := make(map[string]string)
    v6 := make(map[string]uint32)
    var v16 float32 = 1.5
    v17 := v5["hello"]
    for fi18 := 0; fi18 < 1; fi18++ {
        if len(v4) > 0 {
            v4[0] = ((v17 + ((v17 + ("" + v17)) + v17)) + (v17 + v17))
        }
        _ = fi18
    }
    if (false || false) {
        if (false || false) {
            delete(v5, "")
            if false {
                for fi19 := 0; fi19 < 8; fi19++ {
                    switch v17 {
                    case "a":
                        if true {
                            if len(v4) > 0 {
                                v4[0] = ("" + v17)
                            }
                        }
                    case "foo bar":
                        switch (^(v1)) {
                        case 9223372036854775807:
                            var v20 uint = 2
                            _ = v20
                        case -2147483648:
                            fmt.Println(fi19, fi19)
                        case 2147483647:
                            for fi21 := 0; fi21 < 7; fi21++ {
                                var v22 uint64 = 0
                                _ = fi21
                                _ = v22
                            }
                        }
                    case "hello":
                        fmt.Println(v1, v3, v1)
                    }
                    var v23 uint8 = 1
                    _ = fi19
                    _ = v23
                }
            }
            var v24 float32 = (((1.5 * float32(1.5)) * v16) / (3.14159 / 0.1))
            _ = v24
        }
        var v25 float64 = -2.25
        _ = v25
    } else {
        if (float64((v16 / (0.0 + v16))) != -2.25) {
            v26 := GenS1{F0: -2147483648}
            v2 -= (v2 % ((v2) | 1))
            v27 := GenS2{F0: (!((v1 - (v1 >> (((^(v2))) & 63))) != 42)), F1: 4294967295, F2: genFn1(18446744073709551615)}
            _ = v26
            _ = v27
        } else {
            for fi28 := 0; fi28 < 1; fi28++ {
                fmt.Println(v17, v1)
                _ = fi28
            }
            var v29 uint16 = 65535
            _ = v29
        }
        v30 := GenS1{F0: 42}
        for fi31 := 0; fi31 < 18; fi31++ {
            var v32 bool = false
            _ = fi31
            _ = v32
        }
        _ = v30
    }
    v33 := v6[("" + v17)]
    var v34 int32 = -563
    var v35 uint = 18446744073709551615
    var v36 int = 32767
    for fi37 := 0; fi37 < 13; fi37++ {
        v38 := GenS3{F0: 89}
        _ = fi37
        _ = v38
    }
    for fi39 := 0; fi39 < 19; fi39++ {
        v40 := GenS4{F0: (v35 << ((v35) & 63))}
        if false {
            if ((810 < 18446744073709551615) || (!(((v33 != 4294967295) || (0.1 > ((v16 / v16) - (float32(-2.25) - v16)))) || true))) {
                v2 = (v2 << ((18446744073709551615) & 63))
            }
            switch v17 {
            case "a":
                fmt.Println(v34, v2)
            }
            if false {
                _ = len(v4)
                for fi41 := 0; fi41 < 17; fi41++ {
                    switch fi41 {
                    case 409:
                        fmt.Println(v17, v33, fi39)
                    case 0:
                        if (!(42 > 2)) {
                            var v42 string = ((v17 + (v17 + "")) + (("hello" + "hello") + (v17 + "")))
                            fmt.Println(v42)
                            _ = v42
                        }
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi41
                }
                if len(v4) > 0 {
                    v4[0] = v17
                }
            } else {
                var v43 uint = v2
                _ = v43
            }
        } else {
            var v44 float32 = ((v16 / float32(3.14159)) - v16)
            _ = v44
        }
        _ = fi39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
