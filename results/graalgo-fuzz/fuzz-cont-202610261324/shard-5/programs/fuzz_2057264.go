package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
func genFn0() bool {
    v6 := []int{2147483647}
    var v7 uint64 = 2
    _ = v6
    _ = v7
    return (0 != 621)
}

func main() {
    v1 := make(map[string]string)
    var v2 bool = false
    var v3 int = 1
    var v4 float64 = -2.25
    var v5 int8 = -1
    v8 := GenS1{F0: "a", F1: 1}
    for fi9 := 0; fi9 < 19; fi9++ {
        for fi10 := 0; fi10 < 18; fi10++ {
            switch fi10 {
            case 1:
                v11 := GenS1{F0: ("a" + ("\n\t" + "\n\t")), F1: 220}
                _ = v11
            case 4294967295:
                v5 -= (v5 & v5)
            default:
                var v12 int8 = (v5 % ((-128) | 1))
                _ = v12
            }
            if (((v5 & (v5 - (v5 << ((400) & 63)))) == v5) || false) {
                switch "foo bar" {
                case "hello":
                    var v13 int16 = -173
                    v5 = v5
                    var v14 uint8 = 255
                    _ = v13
                    _ = v14
                case "a":
                    delete(v1, (("foo bar" + "a") + ("\n\t" + ("a" + (((((("foo bar" + "foo bar") + "hello") + "a") + ("" + ("" + "hello"))) + "\n\t") + "a")))))
                }
                v3 = fi10
            } else {
                v2 = false
            }
            switch "\n\t" {
            case "\n\t":
                if v2 {
                    var v15 uint8 = 127
                    fmt.Println(v5)
                    _ = v15
                }
                delete(v1, "\n\t")
                v4 = float64((0.0 / (float32((v4 / v4)) * -2.25)))
            case "":
                var v16 uint = 42
                _ = v16
            case "foo bar":
                v17 := GenS1{F0: "", F1: 42}
                _ = v17
            }
            v4 += ((v4 / 1.0) * (v4 * (float64((0.1 - 3.14159)) * (v4 / v4))))
            v18 := []int{fi10, v3, fi9}
            v2 = true
            v19 := GenS1{F0: "a", F1: 663}
            _ = fi10
            _ = v18
            _ = v19
        }
        v1[("\n\t" + "\n\t")] = "\n\t"
        _ = fi9
    }
    var v20 uint32 = 0
    v21 := []int{}
    v20 = 2147483647
    delete(v1, "foo bar")
    var v22 string = "\n\t"
    var v23 string = v22
    var v24 float64 = 1.5
    if len(v21) >= 2 {
        v21 = v21[1:len(v21)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
