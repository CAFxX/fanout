package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 int8; F1 int8; F2 int }
func genFn0() uint16 {
    v5 := GenS1{F0: 225}
    var v6 uint8 = 2
    fmt.Println(v6)
    for fi7 := 0; fi7 < 3; fi7++ {
        for fi8 := 0; fi8 < 10; fi8++ {
            if (65535 <= 266) {
                v6 += 42
                v9 := []int{-9223372036854775808, fi8}
                switch (fi8 ^ (fi7 << ((76) & 63))) {
                case -2147483648:
                    fmt.Println(fi7)
                }
                var v10 int8 = 0
                switch (fi8 & 912) {
                case 32768:
                    v11 := GenS1{F0: (^(v6))}
                    _ = v11
                case 2147483647:
                    v10 += v10
                }
                var v12 int8 = -1
                var v13 bool = true
                _ = v10
                _ = v12
                _ = v13
                _ = v9
            } else {
                if false {
                    v14 := GenS1{F0: 0}
                    fmt.Println(fi8, fi7)
                    v15 := GenS1{F0: (^(v6))}
                    _ = v14
                    _ = v15
                }
                v6 += v6
            }
            if false {
                fmt.Println(v6, fi8)
                var v16 float32 = (0.0 - (3.14159 * (0.0 * 3.14159)))
                var v17 bool = (true || true)
                _ = v16
                _ = v17
            } else {
                var v18 bool = (((((0.0 * 3.14159) + (float64((0.0 - 0.1)) - 0.1)) == ((0.1 - 0.0) - 3.14159)) && ((false || ((("\n\t" + "") + ((("\n\t" + "hello") + "foo bar") + "a")) < "foo bar")) || (943 <= 4294967295))) || false)
                v19 := GenS1{F0: 0}
                if v18 {
                    v6 += 1
                    if (!true) {
                        v20 := []int{fi8, fi7}
                        _ = v20
                    }
                }
                _ = v18
                _ = v19
            }
            v21 := []int{fi7, (fi7 * 0), (^(fi8))}
            _ = fi8
            _ = v21
        }
        _ = fi7
    }
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 int, p1 int64, p2 uint) int {
    for fi22 := 0; fi22 < 11; fi22++ {
        v23 := []int{p0, 4294967295}
        if false {
            if false {
                if len(v23) > 0 {
                    v23[0] = (fi22 * p0)
                }
                var v24 uint = p2
                _ = v24
            }
            p2 += 603
            switch int(p1) {
            case -50:
                p1 += p1
            case -1:
                var v25 uint = 42
                _ = v25
            case 0:
                if (!true) {
                    p0 += (p0 % ((fi22) | 1))
                    switch "foo bar" {
                    case "a":
                        if true {
                            var v26 uint32 = 2
                            v27 := GenS2{F0: "a", F1: p2}
                            for fi28 := 0; fi28 < 12; fi28++ {
                                var v29 uint8 = 42
                                v30 := GenS1{F0: 42}
                                var v31 uint64 = 0
                                _ = fi28
                                _ = v29
                                _ = v30
                                _ = v31
                            }
                            _ = v26
                            _ = v27
                        } else {
                            switch ("hello" + "a") {
                            case "foo bar":
                                var v32 string = "\n\t"
                                _ = v32
                            }
                            var v33 float32 = float32(-2.25)
                            p0 += p0
                            _ = v33
                        }
                    }
                } else {
                    _ = len(v23)
                }
            }
        } else {
            if true {
                p1 -= p1
                var v34 int16 = 1000
                if ((2 == 957) && (false || false)) {
                    var v35 float64 = (-2.25 + 1e10)
                    if true {
                        v34 = 32767
                        for fi36 := 0; fi36 < 20; fi36++ {
                            var v37 uint16 = 255
                            _ = fi36
                            _ = v37
                        }
                        switch -2147483648 {
                        case 9223372036854775807:
                            v34 = (v34 & (v34 * (v34 - (v34 << ((2) & 63)))))
                        case 691:
                            p2 = p2
                        default:
                            p1 = 0
                        }
                    }
                    _ = v35
                } else {
                    v38 := GenS2{F0: "\n\t", F1: p2}
                    var v39 int64 = (-(p1))
                    _ = v38
                    _ = v39
                }
                _ = v34
            } else {
                var v40 string = "a"
                switch fi22 {
                case 128:
                    v40 = ("hello" + v40)
                case 0:
                    for fi41 := 0; fi41 < 19; fi41++ {
                        p2 -= (p2 >> (((p2 >> ((p2) & 63))) & 63))
                        v42 := GenS3{F0: 100, F1: -118, F2: fi41}
                        p1 = p1
                        _ = fi41
                        _ = v42
                    }
                }
                switch (v40 + v40) {
                case "foo bar":
                    p2 = 255
                case "hello":
                    fmt.Println(p1, v40, v40)
                }
                _ = v40
            }
            _ = len(v23)
        }
        _ = fi22
        _ = v23
    }
    var v43 float64 = (1e10 / 1.5)
    _ = p0
    _ = p1
    _ = p2
    _ = v43
    return p0
}

func main() {
    var v1 bool = (true && true)
    var v2 uint8 = 242
    var v3 int64 = 0
    v4 := make(map[string]int)
    var v44 bool = ((42 > 100) || false)
    v2 += v2
    v3 = (v3 | v3)
    v45 := GenS2{F0: "foo bar", F1: 65535}
    fmt.Println(v2, v3, v1)
    fmt.Println(v2, v3)
    v44 = (-1 < 42)
    v46 := GenS3{F0: -127, F1: -1, F2: -2147483648}
    for fi47 := 0; fi47 < 15; fi47++ {
        var v48 float64 = 3.14159
        v49 := []int{}
        var v50 uint16 = 276
        _ = fi47
        _ = v48
        _ = v49
        _ = v50
    }
    v51 := []int{9223372036854775807, 127}
    v51 = append(v51, 2147483647)
    if v44 {
        var v52 uint32 = 42
        _ = v52
    }
    delete(v4, ("a" + "a"))
    var v53 float32 = (-2.25 / 3.14159)
    v3 -= (v3 ^ 9223372036854775807)
    for fi54 := 0; fi54 < 19; fi54++ {
        if v44 {
            var v55 uint64 = 4294967295
            var v56 uint16 = 255
            v57 := v4["foo bar"]
            _ = v55
            _ = v56
            _ = v57
        } else {
            var v58 int32 = -2147483648
            _ = v58
        }
        v59 := v4[("foo bar" + "\n\t")]
        _ = fi54
        _ = v59
    }
    v60 := GenS2{F0: ((("" + ("\n\t" + ("" + "hello"))) + "") + "foo bar"), F1: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v51:", len(v51), v51)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v51
    _ = v53
    _ = v60
}
