package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 string, p1 uint, p2 int32) int32 {
    v4 := GenS1{F0: 65535}
    fmt.Println(p2, p2, p2)
    var v5 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return p2
}

func main() {
    var v1 int = 42
    var v2 int16 = 85
    var v3 uint8 = 255
    v1 -= (-(v1))
    for fi6 := 0; fi6 < 7; fi6++ {
        switch "foo bar" {
        case "hello":
            if false {
                if true {
                    v7 := []int{}
                    fmt.Println(v3, v3, fi6)
                    v1 = v1
                    _ = v7
                } else {
                    v2 -= v2
                    v2 = 42
                    v8 := []int{v1, (fi6 << ((0) & 63)), v1, (v1 * (v1 ^ v1))}
                    _ = v8
                }
                for fi9 := 0; fi9 < 15; fi9++ {
                    switch (("\n\t" + ("a" + (("\n\t" + "\n\t") + "hello"))) + (("\n\t" + ("foo bar" + "")) + "hello")) {
                    case "":
                        v2 = 1000
                    default:
                        v10 := GenS1{F0: 32767}
                        _ = v10
                    }
                    _ = fi9
                }
            }
        }
        switch -578 {
        case -482:
            v11 := []int{v1}
            _ = v11
        default:
            v12 := []int{v1}
            _ = v12
        }
        for fi13 := 0; fi13 < 18; fi13++ {
            var v14 float64 = float64(((0.1 * 1.5) * ((-2.25 * 0.1) / ((3.14159 / (1.5 + 3.14159)) + float32(((float64((0.0 + 1.5)) / 0.1) - (1e10 - 0.1)))))))
            v3 = (v3 >> ((825) & 63))
            _ = fi13
            _ = v14
        }
        _ = fi6
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        switch -136 {
        case -1:
            for fi16 := 0; fi16 < 7; fi16++ {
                var v17 float32 = 0.1
                _ = fi16
                _ = v17
            }
        case 312:
            v3 -= v3
        case 32767:
            for fi18 := 0; fi18 < 14; fi18++ {
                var v19 string = ""
                var v20 int32 = 1
                v1 += fi18
                _ = fi18
                _ = v19
                _ = v20
            }
        }
        for fi21 := 0; fi21 < 10; fi21++ {
            v22 := GenS1{F0: 722}
            _ = fi21
            _ = v22
        }
        if true {
            v2 += (v2 << ((255) & 63))
            for fi23 := 0; fi23 < 3; fi23++ {
                var v24 uint64 = 0
                switch v1 {
                case -266:
                    var v25 string = "foo bar"
                    _ = v25
                case 2147483647:
                    v24 = v24
                case 127:
                    var v26 int32 = -2147483648
                    _ = v26
                }
                var v27 int16 = (v2 * v2)
                _ = fi23
                _ = v24
                _ = v27
            }
        }
        _ = fi15
    }
    var v28 uint8 = v3
    v29 := []int{(v1 % (((v1 << ((787) & 63))) | 1)), (-(v1)), v1, (v1 & v1)}
    var v30 uint64 = 42
    var v31 uint32 = 1
    if false {
        v2 -= v2
        if len(v29) > 0 {
            v29[0] = int(v3)
        }
        var v32 uint64 = (v30 + 0)
        _ = v32
    } else {
        var v33 int16 = -254
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
