package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 42
    v3 := make(map[uint32]int)
    var v4 string = ""
    v5 := make(map[int]uint32)
    v6 := make(map[string]string)
    for fi7 := 0; fi7 < 19; fi7++ {
        if (42 < 2147483647) {
            fmt.Println(v2, fi7)
            var v8 string = v4
            _ = v8
        }
        v9 := GenS1{F0: fi7}
        v4 = v4
        _ = fi7
        _ = v9
    }
    fmt.Println(v4, v2)
    v10 := v6[((v4 + ("" + "")) + "\n\t")]
    if false {
        for fi11 := 0; fi11 < 16; fi11++ {
            v12 := []int{(fi11 & 32768), 558, (fi11 ^ (fi11 << ((18446744073709551615) & 63)))}
            if (!((v2 != v2) && (!(true && false)))) {
                if true {
                    if len(v12) > 0 {
                        v12[0] = (fi11 % ((fi11) | 1))
                    }
                    v1["hello"] = fi11
                    v4 = v10
                } else {
                    for fi13 := 0; fi13 < 1; fi13++ {
                        var v14 float32 = -2.25
                        _ = fi13
                        _ = v14
                    }
                    if len(v12) > 0 {
                        v12[0] = (fi11 * (^(fi11)))
                    }
                    v15 := GenS1{F0: (fi11 << ((1) & 63))}
                    _ = v15
                }
                if (-551 <= -1) {
                    v12 = append(v12, fi11)
                    var v16 uint = 4294967295
                    _ = v16
                }
            } else {
                if false {
                    var v17 uint = 513
                    _ = v17
                } else {
                    var v18 int64 = -1
                    var v19 float64 = 0.0
                    _ = v18
                    _ = v19
                }
                v20 := GenS1{F0: (fi11 / ((4294967295) | 1))}
                switch fi11 {
                case 863:
                    for fi21 := 0; fi21 < 3; fi21++ {
                        fmt.Println(fi11, v2)
                        _ = fi21
                    }
                }
                _ = v20
            }
            _ = fi11
            _ = v12
        }
    }
    v10 = v4
    var v22 int16 = -1
    if (((760 > 42) || true) || false) {
        if true {
            var v23 uint8 = 42
            v23 = 216
            v24 := GenS1{F0: 128}
            _ = v23
            _ = v24
        }
    }
    for fi25 := 0; fi25 < 16; fi25++ {
        var v26 uint16 = v2
        for fi27 := 0; fi27 < 17; fi27++ {
            var v28 uint8 = uint8(v2)
            fmt.Println(v10, v2, fi25)
            _ = fi27
            _ = v28
        }
        if (42 <= int64(v22)) {
            if ((32767 != v2) && (v2 < v26)) {
                v4 = ("foo bar" + ("a" + v4))
                var v29 float64 = -2.25
                switch fi25 {
                case -116:
                    v26 -= (v2 & (^(v26)))
                case 42:
                    var v30 uint16 = (v26 * v2)
                    _ = v30
                }
                _ = v29
            } else {
                var v31 uint64 = 655
                var v32 int64 = 42
                for fi33 := 0; fi33 < 14; fi33++ {
                    if false {
                        switch v10 {
                        case "a":
                            v10 = v4
                            var v34 uint = 42
                            v34 = (v34 >> ((uint(v31)) & 63))
                            _ = v34
                        }
                    }
                    _ = fi33
                }
                _ = v31
                _ = v32
            }
            v1[v10] = 4294967295
        }
        _ = fi25
        _ = v26
    }
    var v35 float32 = 0.1
    v10 = ""
    for fi36 := 0; fi36 < 11; fi36++ {
        var v37 string = v10
        v38 := GenS1{F0: fi36}
        switch v10 {
        case "hello":
            for fi39 := 0; fi39 < 8; fi39++ {
                if (v2 > (^(v2))) {
                    v10 = (v4 + ((v4 + (("\n\t" + "") + v10)) + v37))
                } else {
                    fmt.Println(fi39, v22, fi36)
                    fmt.Println(v22, fi39)
                    var v40 uint32 = 850
                    _ = v40
                }
                v6[("a" + v10)] = v10
                fmt.Println(fi36, v2)
                var v41 string = ("a" + v10)
                _ = fi39
                _ = v41
            }
        }
        _ = fi36
        _ = v37
        _ = v38
    }
    if (!true) {
        v42 := GenS1{F0: -2147483648}
        var v43 int = 0
        _ = v42
        _ = v43
    }
    var v44 uint64 = 42
    switch 42 {
    case 42:
        if true {
            v44 = 907
            if false {
                if (3.14159 != -2.25) {
                    v45 := []int{}
                    var v46 uint = uint(v2)
                    var v47 uint32 = 226
                    _ = v45
                    _ = v46
                    _ = v47
                } else {
                    v10 = (("\n\t" + "foo bar") + "hello")
                    v35 = (0.0 + 0.0)
                }
            } else {
                var v48 int32 = -1
                v49 := GenS1{F0: 42}
                _ = v48
                _ = v49
            }
        }
    default:
        var v50 uint16 = 0
        _ = v50
    }
    var v51 float32 = ((v35 - (v35 * (v35 / (0.1 + v35)))) - ((-2.25 + float32(1e10)) * 0.0))
    v4 = "\n\t"
    var v52 int16 = (v22 / ((v22) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v3
    _ = v35
    _ = v4
    _ = v44
    _ = v5
    _ = v51
    _ = v52
    _ = v6
}
