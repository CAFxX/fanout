package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int }
func genFn0() int {
    var v6 string = "a"
    switch 2147483647 {
    case 277:
        var v7 string = "foo bar"
        _ = v7
    }
    _ = v6
    return -567
}

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 42
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]bool)
    var v5 uint16 = 135
    var v8 int64 = -91
    var v9 float64 = float64((0.0 - (float32(1.5) * 1.5)))
    v8 -= (v8 | 1)
    if true {
        var v10 string = ""
        v11 := GenS1{F0: 42}
        _ = v10
        _ = v11
    } else {
        v12 := GenS1{F0: 33}
        v9 += float64(((0.1 * (0.1 - 3.14159)) / float32((float64((1e10 * (0.1 + 0.0))) - ((v9 * ((1e10 / 3.14159) - 0.0)) * (v9 + 3.14159))))))
        if (("a" <= "") || true) {
            v13 := []int{}
            _ = v13
        } else {
            if (false || false) {
                if false {
                    fmt.Println(v9)
                    for fi14 := 0; fi14 < 11; fi14++ {
                        v15 := GenS1{F0: 983}
                        var v16 int64 = v8
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                    v5 = (v5 % ((808) | 1))
                }
                var v17 string = "\n\t"
                _ = v17
            }
            v18 := GenS1{F0: 580}
            v8 = v8
            _ = v18
        }
        _ = v12
    }
    var v19 uint = 1
    v20 := GenS1{F0: 680}
    v21 := GenS2{F0: -31}
    if (1e10 >= -2.25) {
        var v22 float32 = 1e10
        if false {
            var v23 bool = true
            if (!(v19 <= 201)) {
                var v24 int = genFn0()
                for fi25 := 0; fi25 < 17; fi25++ {
                    if (int8(fi25) > 100) {
                        for fi26 := 0; fi26 < 12; fi26++ {
                            switch (fi26 >> (((v19 + uint(v2))) & 63)) {
                            case 32768:
                                v27 := GenS2{F0: fi25}
                                _ = v27
                            }
                            v28 := v1[("\n\t" + "a")]
                            _ = fi26
                            _ = v28
                        }
                    }
                    v29 := GenS1{F0: 133}
                    v23 = v23
                    _ = fi25
                    _ = v29
                }
                _ = v24
            }
            _ = v23
        } else {
            var v30 int8 = -1
            v30 += v30
            _ = v30
        }
        _ = v22
    } else {
        var v31 bool = false
        v32 := GenS2{F0: 42}
        _ = v31
        _ = v32
    }
    v5 = 2
    for fi33 := 0; fi33 < 8; fi33++ {
        if ((v2 << ((v19) & 63)) <= (v2 | 4294967295)) {
            if false {
                for fi34 := 0; fi34 < 4; fi34++ {
                    var v35 int = int(v5)
                    if (int8(v19) == -57) {
                        fmt.Println(v8, fi33)
                        fmt.Println(v19, v19)
                    }
                    _ = fi34
                    _ = v35
                }
                for fi36 := 0; fi36 < 12; fi36++ {
                    v2 = v2
                    v3[0] = (1 <= 119)
                    if true {
                        fmt.Println(fi33, fi33, v8)
                    }
                    _ = fi36
                }
            }
            v19 = 255
        }
        v8 += (v8 - int64(fi33))
        fmt.Println(v19, fi33)
        _ = fi33
    }
    v4[2] = (uint32(v5) <= 2147483647)
    if true {
        v5 = v5
        for fi37 := 0; fi37 < 11; fi37++ {
            var v38 float32 = 3.14159
            _ = fi37
            _ = v38
        }
        v39 := GenS2{F0: genFn0()}
        _ = v39
    }
    v40 := GenS2{F0: -2147483648}
    switch 32767 {
    case -646:
        if ((0.1 - (float32(v9) + 3.14159)) < float32(v9)) {
            var v41 int8 = 127
            _ = v41
        } else {
            switch ("\n\t" + "") {
            case "":
                switch "a" {
                case "hello":
                    v42 := v1["a"]
                    _ = v42
                case "\n\t":
                    v9 = v9
                case "foo bar":
                    if false {
                        v2 = v2
                        v43 := []int{}
                        _ = v43
                    }
                }
            case "hello":
                var v44 float64 = float64((0.1 + 0.1))
                _ = v44
            }
        }
    case 2147483647:
        v9 = v9
    case 0:
        v45 := []int{42}
        _ = v45
    default:
        var v46 float32 = 3.14159
        _ = v46
    }
    var v47 uint64 = v2
    v48 := v1[("hello" + ((("foo bar" + "") + "a") + "a"))]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v40
    _ = v47
    _ = v48
    _ = v5
    _ = v8
    _ = v9
}
