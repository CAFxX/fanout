package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 uint; F1 int16; F2 int32 }

func main() {
    var v1 int32 = 798
    v2 := make(map[uint32]bool)
    var v3 int16 = -1
    v4 := make(map[string]uint32)
    var v5 int32 = 2147483647
    var v6 string = ("\n\t" + "foo bar")
    var v7 bool = ((467 >= 42) && true)
    if true {
        v8 := GenS1{F0: ((1 < 561) || (!((42 == 2) || (true && ((float32((-2.25 / -2.25)) - 3.14159) <= 0.1)))))}
        for fi9 := 0; fi9 < 7; fi9++ {
            v1 = v1
            _ = fi9
        }
        fmt.Println(v6, v1)
        _ = v8
    } else {
        for fi10 := 0; fi10 < 3; fi10++ {
            for fi11 := 0; fi11 < 4; fi11++ {
                v4[v6] = 2147483648
                v5 = (v1 ^ v5)
                _ = fi11
            }
            v12 := []int{(fi10 & (^(fi10))), -698}
            _ = fi10
            _ = v12
        }
        for fi13 := 0; fi13 < 15; fi13++ {
            var v14 uint = 858
            _ = fi13
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        var v16 int64 = 131
        var v17 bool = v7
        for fi18 := 0; fi18 < 17; fi18++ {
            switch ((("foo bar" + (v6 + v6)) + v6) + v6) {
            case "hello":
                if (!(!v7)) {
                    var v19 int16 = (v3 - v3)
                    if (!true) {
                        if ((v17 || (67 == -128)) || (false && (-792 <= (v1 << ((42) & 63))))) {
                            v3 = v3
                            var v20 uint64 = 300
                            _ = v20
                        } else {
                            if v17 {
                                v3 = v3
                                var v21 float32 = float32((float64(0.1) * -2.25))
                                _ = v21
                            }
                            v22 := []int{1, fi15, (fi18 << ((18446744073709551615) & 63))}
                            _ = v22
                        }
                        v5 = -801
                        if v17 {
                            var v23 uint = 0
                            v24 := []int{fi15}
                            _ = v23
                            _ = v24
                        }
                    } else {
                        v25 := []int{fi18}
                        if (v7 && (!v17)) {
                            v6 = v6
                            if len(v25) >= 2 {
                                v25 = v25[1:len(v25)]
                            }
                            v16 = v16
                        }
                        _ = v25
                    }
                    var v26 float64 = float64((3.14159 * (3.14159 * float32((0.1 + ((float64((1.5 + 0.0)) / float64(3.14159)) - ((3.14159 * 0.1) / (1.5 - (0.1 / 1.5)))))))))
                    _ = v19
                    _ = v26
                }
            }
            _ = fi18
        }
        _ = fi15
        _ = v16
        _ = v17
    }
    v27 := []int{1}
    v1 -= (v5 & (^(v1)))
    if false {
        if (true && (!((v6 + v6) < (v6 + "hello")))) {
            switch (v6 + "") {
            case "hello":
                switch -215 {
                case 2147483647:
                    if v7 {
                        v28 := GenS2{F0: (!v7), F1: (v7 && false)}
                        if ((v7 || ((2 > 213) || (-98 >= 1))) && (881 == 218)) {
                            v6 = ""
                        } else {
                            if (v3 <= v3) {
                                for fi29 := 0; fi29 < 6; fi29++ {
                                    v30 := GenS1{F0: ((v5 > v1) && (!false))}
                                    _ = fi29
                                    _ = v30
                                }
                            }
                        }
                        var v31 int8 = 79
                        _ = v28
                        _ = v31
                    }
                case -1:
                    v4["\n\t"] = 2
                case 1:
                    if ("hello" == v6) {
                        fmt.Println(v5, v1)
                    }
                }
            case "foo bar":
                for fi32 := 0; fi32 < 15; fi32++ {
                    v3 = 170
                    _ = fi32
                }
            }
            var v33 uint32 = 678
            v34 := GenS3{F0: -710, F1: (^(v3))}
            _ = v33
            _ = v34
        }
        if (true || v7) {
            var v35 string = (v6 + "")
            _ = v35
        } else {
            if (57 < 0) {
                var v36 int32 = (v1 + v1)
                if (v7 && (837 == 860)) {
                    var v37 int16 = v3
                    v7 = false
                    _ = v37
                }
                v38 := GenS2{F0: true, F1: v7}
                _ = v36
                _ = v38
            }
            v39 := GenS4{F0: 417, F1: v3, F2: -1}
            _ = v39
        }
    }
    var v40 int64 = -9223372036854775808
    v41 := GenS3{F0: (^(v3)), F1: v3}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
    _ = v7
}
