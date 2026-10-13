package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 uint64 }
type GenS2 struct { F0 uint; F1 int32; F2 uint32 }
type GenS3 struct { F0 int }

func main() {
    var v1 float64 = float64(((0.1 / 0.1) + 3.14159))
    var v2 int32 = 355
    var v3 bool = false
    var v4 int = 4294967295
    var v5 string = "a"
    var v6 uint64 = 809
    switch (("\n\t" + (v5 + v5)) + "foo bar") {
    case "foo bar":
        v7 := GenS1{F0: 2147483647, F1: (v2 >> ((143) & 63)), F2: (v6 | 18446744073709551615)}
        _ = v7
    case "\n\t":
        for fi8 := 0; fi8 < 1; fi8++ {
            v9 := GenS2{F0: 0, F1: -65, F2: 440}
            var v10 float64 = (float64(3.14159) * v1)
            for fi11 := 0; fi11 < 15; fi11++ {
                var v12 int16 = -1
                v2 -= -2147483648
                if (!(223 == 2)) {
                    v10 = float64(0.1)
                    v6 -= 4294967295
                    var v13 uint8 = 127
                    _ = v13
                } else {
                    switch (v5 + (v5 + (v5 + ((("\n\t" + "hello") + ("\n\t" + v5)) + v5)))) {
                    case "a":
                        if (float32(float64(0.0)) == ((3.14159 / 0.1) + ((0.0 * 3.14159) - 1e10))) {
                            v14 := []int{}
                            _ = v14
                        }
                    case "hello":
                        if v3 {
                            if (!(!(!(!((v6 + (v6 * v6)) >= 42))))) {
                                v2 += (-(v2))
                                if true {
                                    v12 = (-(v12))
                                } else {
                                    fmt.Println(fi8)
                                }
                                var v15 int8 = -55
                                _ = v15
                            } else {
                                fmt.Println(fi8, v6, fi11)
                                if (42 != 42) {
                                    var v16 int32 = (v2 | -1)
                                    v1 = 1.5
                                    var v17 float32 = 3.14159
                                    v18 := GenS3{F0: (fi11 & 32767)}
                                    _ = v16
                                    _ = v17
                                    _ = v18
                                } else {
                                    var v19 uint64 = (v6 + (v6 ^ (v6 + (v6 * (v6 % ((v6) | 1))))))
                                    if v3 {
                                        v20 := []int{-750}
                                        v21 := GenS3{F0: (v4 + fi8)}
                                        _ = v20
                                        _ = v21
                                    }
                                    _ = v19
                                }
                                v3 = v3
                            }
                            v6 = 276
                            v22 := GenS1{F0: 65535, F1: (v2 + v2), F2: (^(v6))}
                            _ = v22
                        }
                    default:
                        v23 := []int{2147483647, (fi11 | int(v6)), (v4 * v4)}
                        _ = v23
                    }
                    var v24 int8 = -11
                    _ = v24
                }
                _ = fi11
                _ = v12
            }
            _ = fi8
            _ = v10
            _ = v9
        }
    case "a":
        v4 = (v4 % ((v4) | 1))
    }
    v25 := GenS2{F0: 172, F1: (-(v2)), F2: 350}
    v26 := []int{v4, -23, v4}
    v27 := GenS3{F0: int(v2)}
    var v28 uint32 = 42
    if v3 {
        if false {
            for fi29 := 0; fi29 < 11; fi29++ {
                fmt.Println(v2, v3)
                var v30 float64 = v1
                for fi31 := 0; fi31 < 8; fi31++ {
                    if len(v26) > 0 {
                        v26[0] = (fi31 << ((42) & 63))
                    }
                    var v32 int16 = -1
                    _ = fi31
                    _ = v32
                }
                _ = fi29
                _ = v30
            }
        } else {
            var v33 float32 = 1e10
            v34 := GenS1{F0: (v28 ^ v28), F1: (v2 ^ v2), F2: 4294967295}
            _ = v33
            _ = v34
        }
    } else {
        v4 = 32768
        v3 = ((v6 | 949) > v6)
        if v3 {
            if len(v26) >= 2 {
                v26 = v26[1:len(v26)]
            }
        }
    }
    var v35 float32 = 0.1
    var v36 float32 = float32((v1 * v1))
    v37 := GenS3{F0: (^(v4))}
    v38 := GenS1{F0: 2147483647, F1: v2, F2: (^(v6))}
    v5 = ((("a" + v5) + "\n\t") + "foo bar")
    fmt.Println(v2)
    if false {
        var v39 float32 = float32(1e10)
        switch (v4 ^ 344) {
        case -784:
            for fi40 := 0; fi40 < 13; fi40++ {
                v26 = append(v26, -329)
                var v41 uint8 = 10
                _ = fi40
                _ = v41
            }
        case -81:
            v36 += v39
        }
        v6 = 0
        _ = v39
    } else {
        v36 -= v35
        var v42 float64 = 0.1
        _ = v42
    }
    var v43 uint64 = v6
    switch ((v5 + ("hello" + "")) + ((v5 + "foo bar") + v5)) {
    case "":
        if len(v26) > 0 {
            v26[0] = v4
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
