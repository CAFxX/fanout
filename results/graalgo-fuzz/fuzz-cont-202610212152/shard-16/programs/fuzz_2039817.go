package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 uint64) uint16 {
    p0 += (^(p0))
    v4 := []int{0}
    p0 = 42
    for fi5 := 0; fi5 < 20; fi5++ {
        var v6 string = ""
        var v7 float32 = 0.0
        var v8 int16 = 1000
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = v4
    return 42
}

func main() {
    var v1 string = ""
    v2 := make(map[uint32]uint32)
    v3 := []int{42, 42}
    v9 := v2[65535]
    if ("a" > ((v1 + ("hello" + ("a" + v1))) + v1)) {
        v10 := v2[0]
        if (true && ((!(2 < (v10 + (v10 - v10)))) && true)) {
            v10 = 2
        } else {
            switch (v1 + ((v1 + (v1 + (v1 + (("\n\t" + v1) + (v1 + "hello"))))) + ("a" + v1))) {
            case "hello":
                if (9223372036854775807 == 529) {
                    switch ("\n\t" + "") {
                    case "a":
                        v1 = v1
                    }
                    switch (v1 + v1) {
                    case "hello":
                        for fi11 := 0; fi11 < 2; fi11++ {
                            var v12 float32 = (1.5 - (0.0 + ((1.5 - -2.25) * (float32(((0.1 + (3.14159 / 1.5)) - 1e10)) + 0.0))))
                            for fi13 := 0; fi13 < 12; fi13++ {
                                fmt.Println(fi11, v10)
                                _ = fi13
                            }
                            _ = fi11
                            _ = v12
                        }
                    case "":
                        switch 128 {
                        case 9223372036854775807:
                            for fi14 := 0; fi14 < 13; fi14++ {
                                _ = len(v3)
                                _ = fi14
                            }
                        }
                    default:
                        var v15 uint64 = 4294967295
                        _ = v15
                    }
                    v9 = 2
                } else {
                    var v16 string = ((v1 + "\n\t") + "\n\t")
                    var v17 uint64 = 2
                    _ = v16
                    _ = v17
                }
            case "foo bar":
                v10 += v10
            default:
                v18 := GenS1{F0: 2147483648}
                _ = v18
            }
            var v19 int8 = -128
            _ = v19
        }
        _ = v10
    }
    var v20 uint = 0
    delete(v2, v9)
    v20 = uint(v9)
    for fi21 := 0; fi21 < 6; fi21++ {
        var v22 int = (fi21 % ((127) | 1))
        var v23 float64 = 0.0
        _ = fi21
        _ = v22
        _ = v23
    }
    v24 := v2[v9]
    for fi25 := 0; fi25 < 12; fi25++ {
        var v26 int32 = -910
        var v27 bool = ((!true) || true)
        _ = fi25
        _ = v26
        _ = v27
    }
    fmt.Println(v1)
    v24 += v24
    var v28 int64 = 0
    switch 32768 {
    case 127:
        var v29 bool = false
        _ = v29
    case 0:
        v3 = append(v3, -9223372036854775808)
    }
    var v30 uint = v20
    fmt.Println(v1, v28, v24)
    var v31 float32 = float32(float64(((-2.25 / 1.0) + ((3.14159 - float32(1.5)) * (1.5 / float32(3.14159))))))
    v20 += v30
    v32 := GenS1{F0: (v9 << (((v20 >> ((524) & 63))) & 63))}
    var v33 uint16 = genFn0(18446744073709551615)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v9
}
