package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 uint8; F2 bool }
type GenS3 struct { F0 int; F1 uint64; F2 int32 }
type GenS4 struct { F0 uint16; F1 int64; F2 int }
type GenS5 struct { F0 uint; F1 string; F2 int8 }
func genFn0(p0 string, p1 uint, p2 uint8) int32 {
    v6 := GenS1{F0: 326}
    fmt.Println(p0)
    for fi7 := 0; fi7 < 17; fi7++ {
        if false {
            fmt.Println(p0, fi7, p1)
        }
        fmt.Println(p1, p0)
        v8 := []int{(fi7 + (fi7 - (fi7 - (fi7 << ((p1) & 63))))), (fi7 & 74), fi7}
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 901}
    v10 := GenS1{F0: 517}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v6
    _ = v9
    return 42
}
func genFn1(p0 uint64) uint16 {
    for fi11 := 0; fi11 < 11; fi11++ {
        v12 := GenS2{F0: (("hello" + "") + "a"), F1: 116, F2: (910 == 5)}
        _ = fi11
        _ = v12
    }
    v13 := []int{32767, 733, 92, -9223372036854775808}
    var v14 string = "hello"
    if false {
        v15 := GenS1{F0: 2}
        var v16 string = v14
        _ = v15
        _ = v16
    }
    _ = p0
    _ = v13
    _ = v14
    return 2
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[string]string)
    var v3 int8 = 100
    v4 := make(map[string]string)
    var v5 string = (("foo bar" + ("a" + (("foo bar" + "hello") + ("hello" + "hello")))) + "hello")
    var v17 uint = 0
    var v18 int = 665
    for fi19 := 0; fi19 < 5; fi19++ {
        if true {
            v20 := []int{fi19, 32767, fi19, 1}
            v21 := GenS2{F0: "a", F1: 255, F2: ((1e10 < 3.14159) || (!(18446744073709551615 != 853)))}
            for fi22 := 0; fi22 < 19; fi22++ {
                var v23 uint16 = 600
                var v24 bool = (1 != 4294967295)
                _ = fi22
                _ = v23
                _ = v24
            }
            _ = v20
            _ = v21
        }
        _ = fi19
    }
    for fi25 := 0; fi25 < 19; fi25++ {
        if true {
            v17 += v17
            v1 = -1
            if true {
                if ((!((42 <= -1) && true)) && (42 <= genFn1(18446744073709551615))) {
                    fmt.Println(v1)
                    for fi26 := 0; fi26 < 19; fi26++ {
                        var v27 uint32 = 1
                        delete(v2, "a")
                        _ = fi26
                        _ = v27
                    }
                    fmt.Println(v17, fi25)
                }
                v28 := GenS1{F0: 42}
                _ = v28
            }
        } else {
            v29 := []int{fi25, fi25, (fi25 & v18)}
            v1 = (v1 >> (((v17 * (v17 % ((v17) | 1)))) & 63))
            _ = v29
        }
        v30 := v4[((v5 + (("\n\t" + (v5 + v5)) + v5)) + (v5 + v5))]
        _ = fi25
        _ = v30
    }
    v31 := GenS2{F0: v5, F1: 255, F2: (!false)}
    var v32 uint32 = 29
    fmt.Println(v3)
    v17 += (v17 + 293)
    var v33 string = v5
    v1 += 202
    switch v5 {
    case "":
        v18 += int(v3)
    case "\n\t":
        var v34 int64 = (v1 / ((-9223372036854775808) | 1))
        _ = v34
    }
    if (-1 > -148) {
        if (!false) {
            v3 -= (v3 << ((42) & 63))
            v35 := GenS3{F0: v18, F1: 2, F2: 0}
            _ = v35
        } else {
            switch (v18 >> ((522) & 63)) {
            case 2147483647:
                var v36 string = ((("foo bar" + (v33 + "\n\t")) + (v33 + v33)) + (((v5 + "hello") + "foo bar") + ""))
                _ = v36
            case -9223372036854775808:
                switch v33 {
                case "a":
                    v17 = (v17 - (^(v17)))
                case "foo bar":
                    v32 = v32
                case "\n\t":
                    switch int(v3) {
                    case 492:
                        fmt.Println(v33, v1)
                    }
                default:
                    var v37 bool = (-832 <= (v18 & (v18 + v18)))
                    _ = v37
                }
            }
        }
    }
    var v38 string = v33
    if false {
        fmt.Println(v1, v33)
        fmt.Println(v18, v1, v33)
    } else {
        v3 -= (-(v3))
        fmt.Println(v1)
        for fi39 := 0; fi39 < 13; fi39++ {
            v40 := GenS4{F0: 65535, F1: 9223372036854775807, F2: 639}
            v41 := GenS5{F0: 18446744073709551615, F1: "foo bar", F2: v3}
            _ = fi39
            _ = v40
            _ = v41
        }
    }
    v42 := GenS1{F0: 1}
    v38 = ((((v33 + (v33 + v33)) + ((v38 + "hello") + v5)) + v33) + (v38 + v33))
    v43 := GenS1{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v38
    _ = v4
    _ = v42
    _ = v43
    _ = v5
}
