package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int) int16 {
    var v5 bool = true
    var v6 uint16 = 0
    _ = p0
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 bool = false
    var v2 uint64 = 2
    var v3 uint16 = 1
    var v4 int32 = 2147483647
    v7 := GenS1{F0: 100}
    v2 -= v2
    var v8 uint = 65535
    fmt.Println(v1)
    var v9 uint16 = 42
    var v10 bool = (true && (!false))
    v10 = (((0.1 / 1.0) - 0.0) != ((1.5 - 3.14159) - 3.14159))
    v3 = 493
    v1 = v1
    if true {
        v11 := GenS2{F0: "a"}
        for fi12 := 0; fi12 < 7; fi12++ {
            v13 := []int{fi12, -951, (fi12 << (((v8 | (^(v8)))) & 63)), (fi12 * fi12)}
            for fi14 := 0; fi14 < 11; fi14++ {
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
                v3 = (v3 | v9)
                var v15 float64 = (((0.0 - 1e10) + 1.5) + 1.5)
                _ = fi14
                _ = v15
            }
            fmt.Println(v3, v10)
            _ = fi12
            _ = v13
        }
        _ = v11
    } else {
        v16 := GenS1{F0: 36}
        switch "" {
        case "\n\t":
            var v17 uint64 = v2
            _ = v17
        case "":
            if (false || v10) {
                v18 := GenS2{F0: ("hello" + "foo bar")}
                if false {
                    var v19 bool = (v8 == (^(v8)))
                    v20 := []int{32767, 2147483647}
                    _ = v19
                    _ = v20
                }
                v21 := []int{32768}
                _ = v18
                _ = v21
            } else {
                var v22 uint64 = v2
                _ = v22
            }
        }
        _ = v16
    }
    v8 = v8
    fmt.Println(v10, v8)
    v2 = (^(v2))
    fmt.Println(v2)
    v4 = int32(v8)
    if true {
        var v23 uint64 = 0
        _ = v23
    }
    switch ("\n\t" + "") {
    case "\n\t":
        var v24 uint64 = uint64(v9)
        _ = v24
    case "a":
        v2 -= (v2 ^ (v2 % ((v2) | 1)))
    default:
        var v25 int32 = 1
        _ = v25
    }
    switch int(v8) {
    case 0:
        if v10 {
            if false {
                for fi26 := 0; fi26 < 13; fi26++ {
                    var v27 uint8 = 0
                    _ = fi26
                    _ = v27
                }
                v28 := GenS2{F0: "foo bar"}
                _ = v28
            } else {
                var v29 bool = (127 != 1)
                for fi30 := 0; fi30 < 7; fi30++ {
                    switch (fi30 % (((^(fi30))) | 1)) {
                    case -454:
                        for fi31 := 0; fi31 < 9; fi31++ {
                            v29 = true
                            fmt.Println(v8, v4)
                            if (-2147483648 != fi30) {
                                v32 := []int{int(v4), int(v2), (fi30 / ((360) | 1))}
                                _ = v32
                            }
                            v33 := GenS1{F0: 42}
                            _ = fi31
                            _ = v33
                        }
                    }
                    _ = fi30
                }
                v34 := GenS2{F0: "foo bar"}
                _ = v29
                _ = v34
            }
        } else {
            var v35 bool = (("" + "") != "hello")
            if v35 {
                v3 -= v9
                if v35 {
                    for fi36 := 0; fi36 < 19; fi36++ {
                        if v10 {
                            v37 := []int{(fi36 - fi36), fi36}
                            v35 = v10
                            _ = v37
                        } else {
                            var v38 uint16 = (v9 | v3)
                            _ = v38
                        }
                        switch (fi36 + -775) {
                        case 672:
                            v39 := []int{}
                            _ = v39
                        default:
                            var v40 int64 = -9223372036854775808
                            _ = v40
                        }
                        _ = fi36
                    }
                    v4 += -2147483648
                    v9 = 65535
                } else {
                    v9 = v9
                    v41 := []int{1, 127, 32768}
                    fmt.Println(v3, v3, v2)
                    _ = v41
                }
                fmt.Println(v2, v1)
            }
            _ = v35
        }
    case 32768:
        fmt.Println(v2)
    case -282:
        v42 := GenS3{F0: 127}
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
