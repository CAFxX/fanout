package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0() uint {
    for fi4 := 0; fi4 < 7; fi4++ {
        v5 := GenS1{F0: 170}
        v6 := GenS1{F0: -11}
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 float64 = 0.1
    _ = v7
    return 45
}
func genFn1(p0 int8, p1 int16) int32 {
    var v8 int32 = 1
    var v9 string = "\n\t"
    for fi10 := 0; fi10 < 1; fi10++ {
        var v11 int8 = p0
        v12 := []int{(fi10 | fi10)}
        _ = fi10
        _ = v11
        _ = v12
    }
    fmt.Println(p0, p1, v8)
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return -53
}

func main() {
    var v1 int64 = 1
    var v2 uint32 = 4294967295
    var v3 uint64 = 2
    switch "" {
    case "a":
        fmt.Println(v2, v1)
    default:
        switch 9223372036854775807 {
        case 2147483647:
            v3 = v3
        }
    }
    for fi13 := 0; fi13 < 3; fi13++ {
        v14 := []int{fi13}
        v15 := GenS1{F0: 877}
        switch (("\n\t" + "hello") + ("hello" + "foo bar")) {
        case "":
            switch (("" + "\n\t") + "hello") {
            case "":
                if len(v14) > 0 {
                    v14[0] = 0
                }
            }
        case "hello":
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
        default:
            if (195 == 42) {
                var v16 uint64 = 0
                var v17 float32 = -2.25
                _ = v16
                _ = v17
            } else {
                var v18 uint64 = (^(v3))
                v14 = append(v14, fi13)
                var v19 string = ((("foo bar" + "a") + ("\n\t" + "\n\t")) + "hello")
                _ = v18
                _ = v19
            }
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    v2 = uint32(v3)
    if false {
        var v20 uint64 = v3
        _ = v20
    }
    v2 += 2147483648
    switch -1 {
    case 128:
        v21 := []int{121, 32768, 782, -922}
        _ = v21
    default:
        fmt.Println(v2)
    }
    fmt.Println(v1, v3, v1)
    v1 -= -1
    var v22 bool = (-567 <= 42)
    fmt.Println(v3, v2, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
}
